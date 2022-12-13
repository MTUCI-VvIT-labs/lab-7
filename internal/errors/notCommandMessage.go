package errors

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func NotCommandMessage(bot *botapi.BotAPI, msg botapi.MessageConfig) {
	msg.Text = "Это не команда. Попробуйте еще раз."
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
