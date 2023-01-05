package handlers

import (
	"sort"
	"strconv"
	"time"
	"todo_app/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func CurrentTasksHandler(bot *tgbotapi.BotAPI, updMessage *tgbotapi.Message, DB *gorm.DB, sortType string) {
	var tasks []db.TaskModel
	rows := DB.Where("chat_id = ? AND status = ?", updMessage.Chat.ID, InProgress).Find(&tasks)
	if rows.RowsAffected == 0 {
		msg := tgbotapi.NewMessage(updMessage.Chat.ID, "У вас нет задач в работе")
		bot.Send(msg)
		return
	}

	listOfTasks := "Чтобы завершить задачу, введите ее номер.\nЕсли не хотите завершать задачу, введите /stop\n\n"
	sort.Slice(tasks, func(i, j int) bool {
		if sortType == "time" {
			return tasks[i].TakenToWorkAt.Before(tasks[j].TakenToWorkAt)
		} else {
			return PriorityAssignLevel(tasks[i].Priority) > PriorityAssignLevel(tasks[j].Priority)
		}
	})

	for index, element := range tasks {
		listOfTasks += strconv.Itoa(index+1) + ") " + element.Name + ".  " + PriorityToClear(element.Priority) + " приоритет.\n"
	}
	msg := tgbotapi.NewMessage(updMessage.Chat.ID, listOfTasks)
	bot.Send(msg)
}

func CompleteTask(bot *tgbotapi.BotAPI, updMessage *tgbotapi.Message, DB *gorm.DB) {
	var tasks []db.TaskModel
	rows := DB.Where("chat_id = ? AND status = ?", updMessage.Chat.ID, InProgress).Find(&tasks)
	if rows.RowsAffected == 0 {
		return
	}

	numberOfTask, err := strconv.Atoi(updMessage.Text)
	if err != nil {
		return
	}

	taskToWork := tasks[numberOfTask-1]
	DB.Model(&db.TaskModel{}).Where("id = ?", taskToWork.ID).Updates(db.TaskModel{Status: Completed, CompletedAt: time.Now()})
	msg := tgbotapi.NewMessage(updMessage.Chat.ID, "Задача '"+taskToWork.Name+"' завершена!")
	bot.Send(msg)
}
