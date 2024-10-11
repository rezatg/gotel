package gotel

import "github.com/rezatg/gotel/types"

func (c *Client) OnMessage(handler func(client *Client, update types.Message)) {
	if c.handlers == nil {
		c.handlers = make(map[string][]any)
	}
	c.handlers["message"] = append(c.handlers["message"], handler)
}
