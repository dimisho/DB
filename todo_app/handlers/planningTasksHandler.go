package handlers

import (
	"strconv"
	"time"
	"todo_app/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func PlanningTasksHandler(bot *tgbotapi.BotAPI, updMessage *tgbotapi.Message, DB *gorm.DB) {
	var tasks []db.TaskModel
	rows := DB.Where("chat_id = ? AND status = ?", updMessage.Chat.ID, Planning).Find(&tasks)
	if rows.RowsAffected == 0 {
		msg := tgbotapi.NewMessage(updMessage.Chat.ID, "У вас нет запланированных задач")
		bot.Send(msg)
		return
	}

	listOfTasks := "Чтобы взять задачу в работу, введите ее номер.\nЕсли не хотите брать задачу, введите /stop\n\n"
	for index, element := range tasks {
		listOfTasks += strconv.Itoa(index+1) + ") " + element.Name + ".  " + PriorityToClear(element.Priority) + " приоритет.\n"
	}
	msg := tgbotapi.NewMessage(updMessage.Chat.ID, listOfTasks)
	bot.Send(msg)
}

func TakeTaskToWork(bot *tgbotapi.BotAPI, updMessage *tgbotapi.Message, DB *gorm.DB) {
	var tasks []db.TaskModel
	rows := DB.Where("chat_id = ? AND status = ?", updMessage.Chat.ID, Planning).Find(&tasks)
	if rows.RowsAffected == 0 {
		return
	}

	numberOfTask, err := strconv.Atoi(updMessage.Text)
	if err != nil {
		return
	}

	taskToWork := tasks[numberOfTask-1]
	DB.Model(&db.TaskModel{}).Where("id = ?", taskToWork.ID).Updates(db.TaskModel{Status: InProgress, TakenToWorkAt: time.Now()})
	msg := tgbotapi.NewMessage(updMessage.Chat.ID, "Задача '"+taskToWork.Name+"' взята в работу!")
	bot.Send(msg)
}
