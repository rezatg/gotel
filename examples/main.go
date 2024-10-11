package main

import (
	"github.com/Laky-64/gologging"
	bot "github.com/rezatg/gotel"
)

func main() {
	// gologging.Debug("Setting up kernel of the protogen")
	// gologging.Info("Protogen started")
	gologging.Warn("Protogen is warming up")
	// gologging.Error("Protogen failed to start")
	// gologging.Fatal("Protogen crashed")

	client, _ := bot.NewBotAPI(bot.Config{
		Token: "5789687015:AAElrhOTDR3T5_PpCljLVmAUpZZ9aiju5bI",
	})

	client.SendMessage(&bot.SendMessageParams{ChatID: 1033512845, Text: "test"})

}
