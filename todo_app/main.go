package main

import (
	"log"
	"os"
	"todo_app/db"
	"todo_app/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func checkEnv() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var globalAction = ""

func main() {
	checkEnv()
	DB := db.InitDB()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			// cron.StartCron(bot, update.Message.Chat.ID, DB)
			if update.Message.IsCommand() {
				if update.Message.Text == "/start" {
					handlers.StartHandler(bot, update.Message, DB)
				} else {
					globalAction = ""
					handlers.DefaultHandler(bot, update.Message)
				}
			} else {
				callDefaultHandler := true
				if globalAction != "" {
					switch globalAction {
					case "create":
						handlers.CreateTaskHandler(bot, update.Message, DB)
						// cron.StartCron(bot, update.Message.Chat.ID, DB)
						globalAction = ""
					case "planning":
						handlers.TakeTaskToWork(bot, update.Message, DB)
						globalAction = ""
					case "current":
						handlers.CompleteTask(bot, update.Message, DB)
						globalAction = ""
					case "stats":
						handlers.StatsHandler(bot, update.Message, DB)
						globalAction = ""
						callDefaultHandler = false
					}
				}
				if callDefaultHandler {
					handlers.DefaultHandler(bot, update.Message)
				}
			}
		} else if update.CallbackQuery != nil {
			// cron.StartCron(bot, update.CallbackQuery.Message.Chat.ID, DB)
			switch update.CallbackQuery.Data {
			case "/create":
				handlers.GetNewTaskInfo(bot, update.CallbackQuery.Message)
				globalAction = "create"
			case "/currentByTime":
				handlers.CurrentTasksHandler(bot, update.CallbackQuery.Message, DB, "time")
				globalAction = "current"
			case "/currentByPriority":
				handlers.CurrentTasksHandler(bot, update.CallbackQuery.Message, DB, "priority")
				globalAction = "current"
			case "/planning":
				handlers.PlanningTasksHandler(bot, update.CallbackQuery.Message, DB)
				globalAction = "planning"
			case "/completed":
				handlers.CompletedTasksHandler(bot, update.CallbackQuery.Message, DB)
			case "/stats":
				handlers.GetPeriodStats(bot, update.CallbackQuery.Message, DB)
				globalAction = "stats"
			default:
				handlers.DefaultHandler(bot, update.CallbackQuery.Message)
			}
		}
	}
}
