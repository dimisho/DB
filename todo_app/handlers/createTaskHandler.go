package handlers

import (
	"fmt"
	"strings"
	"time"
	"todo_app/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func priorityToConst(priority string) string {
	switch priority {
	case "Низкий":
		return Low
	case "Средний":
		return Medium
	case "Высокий":
		return High
	default:
		return Low
	}
}
func GetNewTaskInfo(bot *tgbotapi.BotAPI, updMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(updMessage.Chat.ID, "Введите, разделяя точкой описание задачи, приоритет (низкий, средний, "+
		"высокий), дату и время когда уведомить о задаче.\n\nНапример: Встретить курьера. Высокий. 08-01-2023 8:00\n\n"+
		"Введите /stop, если не хотите создавать задачу")
	bot.Send(msg)
}

func CreateTaskHandler(bot *tgbotapi.BotAPI, updMessage *tgbotapi.Message, DB *gorm.DB) {
	values := strings.Split(updMessage.Text, ". ")
	notificationDate, err := time.ParseInLocation("02-01-2006 15:04", values[2], time.Local)
	if err == nil {
		fmt.Println(notificationDate)
	}
	newTask := db.TaskModel{ChatID: uint64(updMessage.Chat.ID), Name: values[0], Priority: priorityToConst(values[1]),
		Status: Planning, Notification: notificationDate}
	DB.Create(&newTask)
	msg := tgbotapi.NewMessage(updMessage.Chat.ID, "Задача создана! Чтобы взять ее в работу, перейдите в 'Запланированные задачи'")
	bot.Send(msg)
}
