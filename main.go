package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var token = "1090407046:AAFittPvxi8ZhtDgAu58KDjv0FuJqH66r8I"
var apiUrl = "https://api.telegram.org/bot" + token + "/"

func main(){
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for {
		select {
		case update := <-updates:
			// Пользователь, который написал боту
			UserName := update.Message.From.UserName

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID

			// Текст сообщения
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// Ответим пользователю его же сообщением
			reply := Text

			if Text == "/connect" {
				// Созадаем сообщение
				msg := tgbotapi.NewMessage(ChatID, "Хочешь познакомиться?")
				bot.Send(msg)

			} else {
				// Созадаем сообщение
				msg := tgbotapi.NewMessage(ChatID, reply)
				bot.Send(msg)
			}
		}
	}
}
