package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	Planning   = "PLANNING"
	InProgress = "IN_PROGRESS"
	Completed  = "COMPLETED"
)

const (
	Low    = "LOW"
	Medium = "MEDIUM"
	High   = "HIGH"
)

func PriorityAssignLevel(priority string) int {
	switch priority {
	case Low:
		return 1
	case Medium:
		return 2
	case High:
		return 3
	default:
		return 0
	}
}

func PriorityToClear(priority string) string {
	switch priority {
	case Low:
		return "Низкий"
	case Medium:
		return "Средний"
	case High:
		return "Высокий"
	default:
		return "Низкий"
	}
}

func StatusToClear(status string) string {
	switch status {
	case Planning:
		return "Запланированно"
	case InProgress:
		return "В работе"
	default:
		return "Завершено"
	}
}

var defaultKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Создать задачу", "/create"),
		tgbotapi.NewInlineKeyboardButtonData("Запланированные", "/planning"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Текущие по времени", "/currentByTime"),
		tgbotapi.NewInlineKeyboardButtonData("Текущие по приоритету", "/currentByPriority"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Посмотреть статистику", "/stats"),
		tgbotapi.NewInlineKeyboardButtonData("Завершенные", "/completed"),
	),
)

func DefaultHandler(bot *tgbotapi.BotAPI, updMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(updMessage.Chat.ID, "Выберите действие")
	msg.ReplyMarkup = defaultKeyboard
	bot.Send(msg)
}
