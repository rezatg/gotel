package gotel

import "github.com/rezatg/gotel/types"

func (c *Client) SendMessage(parameters *SendMessageParams) (*types.Message, error) {
	// switch parameters.ChatID.(type) {
	// case int64:
	// case string:
	// 	parameters.ChatID = parameters.ChatID.(int64)

	// case *Message:
	// 	parameters.ChatID = parameters.ChatID.(*Message).Chat.ID

	// default:
	// 	return nil, nil //errors.New("Chat ID is wrong.")
	// }

	var message *types.Message
	var err error = c.rawRequest("sendMessage", parameters, &message)

	return message, err
}
