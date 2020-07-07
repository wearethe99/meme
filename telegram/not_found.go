package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"meme/container"
)

func NotFound(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, container.BotText)
	msg.ParseMode = "markdown"

	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}