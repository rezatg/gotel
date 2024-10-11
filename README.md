A Go library for creating telegram bots.

```go
package main

import (
	"github.com/mymmrac/telego"

   bot "github.com/rezatg/gotel"
)

func main() {
	// Bot token from variables
	const token string = "5789687015:AAElrhOTDR3T5_PpCljLVmAUpZZ9aiju5bI"

   client, _ := bot.NewBotAPI(bot.Config{
		Token:   token,
		Offline: true,
	})

	client.SendMessage(&bot.SendMessageParams{ChatID: 1033512845, Text: "test"})
}

```