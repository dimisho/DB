package handlers

import (
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
	"todo_app/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

func GetPeriodStats(bot *tgbotapi.BotAPI, updMessage *tgbotapi.Message, DB *gorm.DB) {
	msg := tgbotapi.NewMessage(updMessage.Chat.ID, "Введите период для вывода статистики через пробел в формате 20-01-2023\n\n"+
		"Например: 08-01-2023 10-01-2023\n\n"+
		"Введите /stop, если не хотите смотреть статистику")
	bot.Send(msg)
}

func StatsHandler(bot *tgbotapi.BotAPI, updMessage *tgbotapi.Message, DB *gorm.DB) {
	msg := tgbotapi.NewMessage(updMessage.Chat.ID, "Чтобы вызвать главное меню, введите что-нибудь. Например /stop")
	bot.Send(msg)

	values := strings.Split(updMessage.Text, " ")

	date1, err := time.Parse("02-01-2006", values[0])
	if err != nil {
		panic(err)
	}
	date2, err := time.Parse("02-01-2006", values[1])
	if err != nil {
		panic(err)
	}
	var tasks []db.TaskModel
	rows := DB.Where("chat_id = ? AND status = ? AND completed_at >= ? AND completed_at <= ?",
		updMessage.Chat.ID, Completed, date1, date2).Find(&tasks)

	numberOfCompletedTasks := int(rows.RowsAffected)

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].CompletedAt.Before(tasks[j].CompletedAt)
	})

	result := "Количество выполненных задач: " + strconv.Itoa(numberOfCompletedTasks) + "\n\n"
	for index, element := range tasks {
		dur := element.CompletedAt.Sub(element.TakenToWorkAt).Round(time.Second)
		log.Printf("++++++++++++++")
		log.Printf(dur.String())
		log.Printf("++++++++++++++")
		result += strconv.Itoa(index+1) + ") " + element.Name + ".  " + dur.String() + "\n"
	}
	msg = tgbotapi.NewMessage(updMessage.Chat.ID, result)
	bot.Send(msg)
}
