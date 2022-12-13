package commands

import (
	"MTUCI-VvIT-labs/lab-7/internal/entities"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

func Start(bot *botapi.BotAPI, msg botapi.MessageConfig, username string) {
	msg.Text = "Привет, " + username + " !"
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}

	time.Sleep(1 * time.Second)

	msg.Text = "Чтобы посмотреть список команд напишите: /help"

	msg.ReplyMarkup = entities.Keyboard

	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
