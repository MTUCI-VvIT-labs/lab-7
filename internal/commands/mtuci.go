package commands

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func MTUCI(bot *botapi.BotAPI, msg botapi.MessageConfig) {
	msg.Text = "Чтобы попасть на официальный сайт МТУСИ, перейдите по ссылке: https://mtuci.ru/"
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
