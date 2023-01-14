package cron

import (
	"sort"
	"time"
	"todo_app/db"
	"todo_app/handlers"

	"github.com/go-co-op/gocron"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

var scheduler gocron.Scheduler

func StartCron(bot *tgbotapi.BotAPI, chatId int64, DB *gorm.DB) {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(1).Minute().Do(notificationHandler, bot, chatId, DB)
	scheduler.StartAsync()
}

func notificationHandler(bot *tgbotapi.BotAPI, chatId int64, DB *gorm.DB) {
	scheduler.Clear()
	var tasks []db.TaskModel
	now := time.Now()
	rows := DB.Where("chat_id = ? AND status != ? AND notification > ?", chatId, handlers.Completed, now).Find(&tasks)
	if rows.RowsAffected == 0 {
		return
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Notification.Before(tasks[j].Notification)
	})

	for _, task := range tasks {
		// scheduler.Every(1).Day().StartAt(task.Notification).Do(sendNotification, bot, chatId, task)
		if task.Notification.Equal(time.Now()) {
			sendNotification(bot, chatId, task)
		}
	}
	// scheduler.StartAsync()

}

func sendNotification(bot *tgbotapi.BotAPI, chatId int64, task db.TaskModel) {
	msg := tgbotapi.NewMessage(chatId, "Уведомление о задаче '"+task.Name+"'!\n\n"+
		handlers.PriorityToClear(task.Priority)+" приоритет. Статус: "+handlers.StatusToClear(task.Status))
	bot.Send(msg)
}
