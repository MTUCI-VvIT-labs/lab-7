package commands

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Help(bot *botapi.BotAPI, msg botapi.MessageConfig) {
	msg.Text = "Я - бот, который поможет узнать Вам расписание. \n" +
		"Вот список моих команд:\n\n" +
		"/week - информация о текущей неделе\n" +
		"/mtuci - информация о МТУСИ\n" +
		"/help - список команд"
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
