package main

import (
	"MTUCI-VvIT-labs/lab-7/internal/router"
	"MTUCI-VvIT-labs/lab-7/pkg/pg"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const botToken = "5776628931:AAEXdAH2bUX53-X9MVULvB1zAaX-aJ1j2xc"
const postgresUrl = "postgres://postgres:123456@localhost/postgres?sslmode=disable"

func main() {
	err := pg.ConnectToDB(postgresUrl)
	if err != nil {
		log.Panic(err)
	}

	bot, err := botapi.NewBotAPI(botToken) // создание бота
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName) // подтверждение авторизации

	u := botapi.NewUpdate(0) // создание конфигурации для получения обновлений
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u) // получение канала обновлений
	for update := range updates {    // обработка обновлений
		if update.Message != nil { // проверка наличия сообщения
			router.Route(bot, update)
		}
	}
}
