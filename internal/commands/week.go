package commands

import (
	"MTUCI-VvIT-labs/lab-7/internal/entities"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func Week(bot *botapi.BotAPI, msg botapi.MessageConfig) {
	week := entities.NewWeek()

	msg.Text = "Сейчас идет неделя №" + strconv.Itoa(week.WeekNumber) + ". Она " + week.WeekType + "."
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}
