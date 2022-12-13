package entities

import botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var Keyboard = botapi.NewReplyKeyboard(
	botapi.NewKeyboardButtonRow(
		botapi.NewKeyboardButton("Понедельник"),
		botapi.NewKeyboardButton("Вторник"),
		botapi.NewKeyboardButton("Среда"),
		botapi.NewKeyboardButton("Четверг")),
	botapi.NewKeyboardButtonRow(
		botapi.NewKeyboardButton("Текущая неделя"),
		botapi.NewKeyboardButton("Пятница"),
		botapi.NewKeyboardButton("Следующая неделя")),
)
