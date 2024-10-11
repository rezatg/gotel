package gotel

import (
	"github.com/goccy/go-json"
)

type Config struct {
	// Bot's unique authentication token
	Token string

	// Updates channel capacity, defaulted to 100.
	updateCap int

	// Telegram Bot API service URL, defaulted to ""
	BaseUrl string

	// Telegram Bot API file URL, defaulted to ""
	BaseFileUrl string

	Workers int

	Poller Poller

	AllowedIP bool

	// Offline allows to create a bot without network for testing purposes.
	Offline bool
}

type Update struct {
}

type Response struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result,omitempty"`
	Description json.RawMessage `json:"description,omitempty"`
}

// will be able to override the default mode by passing a new one.
// ParseMode ParseMode

// Poller is the provider of Updates.
// Poller Poller

// Synchronous prevents handlers from running in parallel.
// It makes ProcessUpdate return after the handler is finished.
// Synchronous bool

// Verbose forces bot to log all upcoming requests.
// Use for debugging purposes only.
// Verbose bool

// ParseMode used to set default parse mode of all sent messages.
// It attaches to every send, edit or whatever method. You also
// will be able to override the default mode by passing a new one.
// ParseMode ParseMode

// OnError is a callback function that will get called on errors
// resulted from the handler. It is used as post-middleware function.
// Notice that context can be nil.
// OnError func(error, Context)

// HTTP Client used to make requests to telegram api
// Client *http.Client

// Offline allows to create a bot without network for testing purposes.

// }

// type Bot struct {
// 	Self  *User
// 	Token string
// 	URL   string
// Updates chan Update
// Poller  Poller
// onError func(error, Context)

// group       *Group
// handlers    map[string]HandlerFunc
// synchronous bool
// verbose     bool
// parseMode   ParseMode
// stop        chan chan struct{}
// client      *http.Client

// stopMu     sync.RWMutex
// stopClient chan struct{}
// }
