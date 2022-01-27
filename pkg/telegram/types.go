package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	Bot      *tgbotapi.BotAPI
	Messages Messages
}

type Messages struct {
	Responses
	Errors
}

const (
	COMMAND_MESSAGE = iota
	UNKNOWN_COMMAND_MESSAGE
	KEYBOARD_MESSAGE
	TEXT_MESSAGE
	TAKS_NAME_MESSAGE
	CALLBACK_MESSAGE
)

const (
	commandStart = "start"
)

var COMMANDS = [...]string{"/start"}
