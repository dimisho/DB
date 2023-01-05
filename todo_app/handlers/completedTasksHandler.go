package handlers

import (
	"strconv"
	"todo_app/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func CompletedTasksHandler(bot *tgbotapi.BotAPI, updMessage *tgbotapi.Message, DB *gorm.DB) {
	var tasks []db.TaskModel
	rows := DB.Where("chat_id = ? AND status = ?", updMessage.Chat.ID, Completed).Find(&tasks)
	if rows.RowsAffected == 0 {
		msg := tgbotapi.NewMessage(updMessage.Chat.ID, "У вас нет завершенных задач")
		bot.Send(msg)
		return
	}

	listOfTasks := "Чтобы вызвать главное меню, введите что-нибудь. Например /stop\n\n"
	for index, element := range tasks {
		listOfTasks += strconv.Itoa(index+1) + ") " + element.Name + ".  " + PriorityToClear(element.Priority) + " приоритет.\n"
	}
	msg := tgbotapi.NewMessage(updMessage.Chat.ID, listOfTasks)
	bot.Send(msg)
}
