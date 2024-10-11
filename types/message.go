package types

type Message struct {
	MessageID int `json:"message_id"`

	From *User `json:"from,omitempty"`

	// Chat *Chat `json:"chat"`

	// ForwardOrigin *MessageOrigin `json:"forward_origin,omitempty"`
	// ChatShared    *ChatShared    `json:"chat_shared,omitempty"`

	Date int64  `json:"date"`
	Text string `json:"text,omitempty"`

	// ReplyToMessage *Message        `json:"reply_to_message,omitempty"`
	// Entities       []MessageEntity `json:"entities,omitempty"`
	// LinkPreviewOptions  *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	HasProtectedContent bool `json:"has_protected_content,omitempty"`

	// Photo     []PhotoSize `json:"photo,omitempty"`
	// Audio     *Audio      `json:"audio,omitempty"`
	// Document  *Document   `json:"document,omitempty"`
	// Video     *Video      `json:"video,omitempty"`
	// Voice     *Voice      `json:"voice,omitempty"`
	// Sticker   *Sticker    `json:"sticker,omitempty"`
	// Animation *Animation  `json:"animation,omitempty"`
	// VideoNote *VideoNote  `json:"video_note,omitempty"`

	Caption string `json:"caption,omitempty"`

	// ReplyMarkup InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
