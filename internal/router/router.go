package router

import (
	"MTUCI-VvIT-labs/lab-7/internal/buttons"
	"MTUCI-VvIT-labs/lab-7/internal/commands"
	"MTUCI-VvIT-labs/lab-7/internal/errors"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var buttonTextList = []string{"Понедельник", "Вторник	", "Среда", "Четверг", "Пятница", "Следующая неделя", "Текущая неделя"}

func Route(bot *botapi.BotAPI, update botapi.Update) {
	msg := botapi.NewMessage(update.Message.Chat.ID, "") // создаем пустое сообщение

	logMessage(update)
	if isButton, button := isButtonPressed(update); isButton {
		buttons.ButtonHandler(bot, msg, button)
	} else if !update.Message.IsCommand() { // если сообщение не является командой, то пользователь получает сообщение об ошибке
		errors.NotCommandMessage(bot, msg)
	} else {
		switch update.Message.Command() {
		case "start":
			commands.Start(bot, msg, update.Message.From.UserName)
		case "week":
			commands.Week(bot, msg)
		case "mtuci":
			commands.MTUCI(bot, msg)
		case "help":
			commands.Help(bot, msg)
		default:
			errors.UnknownCommandMessage(bot, msg)
		}
	}
}

func logMessage(message botapi.Update) {
	log.Printf("[%s] %s", message.Message.From.UserName, message.Message.Text)
}

func isButtonPressed(message botapi.Update) (bool, string) {
	for _, buttonText := range buttonTextList {
		if message.Message.Text == buttonText {
			return true, buttonText
		}
	}
	return false, ""
}
