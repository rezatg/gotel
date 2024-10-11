package main

import (
	"github.com/Laky-64/gologging"
	bot "github.com/rezatg/gotel"
)

const token string = "5789687015:AAElrhOTDR3T5_PpCljLVmAUpZZ9aiju5bI"

func main() {
	// gologging.Debug("Setting up kernel of the protogen")
	// gologging.Info("Protogen started")
	gologging.Warn("Protogen is warming up")
	// gologging.Error("Protogen failed to start")
	// gologging.Fatal("Protogen crashed")

	client, _ := bot.NewBotAPI(bot.Config{
		Token:   token,
		Offline: true,
	})

	client.SendMessage(&bot.SendMessageParams{ChatID: 1033512845, Text: "test"})

}
