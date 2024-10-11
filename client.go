package gotel

import (
	"errors"
	"regexp"
	"sync"

	"github.com/rezatg/gotel/types"
)

const (
	defaultApiURL     string = "https://api.telegram.org"
	defaultApiFileURL string = "https://api.telegram.org/file/bot"

	tokenRegexp string = `^\d+:[\w-]{35}$`

	defaultUpdatesChanCap int = 1024
	defaultWorkers        int = 1
)

// New creates new Bot instance
func NewBotAPI(s Config) (*Client, error) {
	if !ValidateToken(s.Token) {
		return nil, errors.New("you must pass the token you received from @Botfather")
	}

	if s.updateCap == 0 {
		s.updateCap = 100
	}

	if s.BaseUrl == "" {
		s.BaseUrl = defaultApiURL
	}

	if s.BaseFileUrl == "" {
		s.BaseFileUrl = defaultApiFileURL
	}

	if s.Workers == 0 {
		s.Workers = defaultWorkers
	}

	var bot *Client = &Client{
		token: s.Token,

		baseUrl:     s.BaseUrl + s.Token,
		baseFileUrl: s.BaseFileUrl + s.Token,

		Poller: FastHTTPCaller{},

		updates: make(chan *Update, s.updateCap),
	}

	if s.Offline {
		bot.self = new(types.User)
	}

	return bot, nil
}

// Bot represents Telegram Bot main object
type Client struct {
	token string

	baseUrl     string
	baseFileUrl string

	self *types.User

	Poller Poller

	workers int
	updates chan *Update

	handlers   map[string][]any
	handlersMx sync.RWMutex

	isRunning bool

	allowedUpdates []string
}

// validateToken validates if token matches format
func ValidateToken(token string) bool {
	return regexp.MustCompile(tokenRegexp).MatchString(token)
}
