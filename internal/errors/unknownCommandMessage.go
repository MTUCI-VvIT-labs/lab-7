package errors

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func UnknownCommandMessage(bot *botapi.BotAPI, msg botapi.MessageConfig) {
	msg.Text = "Извините, я Вас не понял."
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
