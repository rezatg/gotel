package gotel

type SendMessageParams struct {
	ChatID any    `json:"chat_id"`
	Text   string `json:"text"`
}
