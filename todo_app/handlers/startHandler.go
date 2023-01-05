package handlers

import (
	"todo_app/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func StartHandler(bot *tgbotapi.BotAPI, updMessage *tgbotapi.Message, DB *gorm.DB) {
	var userModel db.UserModel
	exist := DB.Find(&userModel, "chat_id = ?", updMessage.Chat.ID)
	if exist.RowsAffected < 1 {
		newUser := db.UserModel{ChatID: uint64(updMessage.Chat.ID), Username: updMessage.From.UserName}
		DB.Create(&newUser)
	}

	msg := tgbotapi.NewMessage(updMessage.Chat.ID, "Привет, "+updMessage.From.FirstName+
		"!\n <Описание бота, что он умеет>\n Введите что-нибудь, чтобы начать.")
	bot.Send(msg)
}
