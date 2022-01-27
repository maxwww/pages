package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/maxwww/pages/pkg/telegram"

	"github.com/maxwww/pages/config"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	botApi, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatalf("error start bot api: %s", err.Error())
	}
	//botApi.Debug = true

	bot := telegram.NewBot(botApi, cfg.Messages)

	if err := bot.Start(); err != nil {
		log.Fatalf("error start bot: %s", err.Error())
	}
}
