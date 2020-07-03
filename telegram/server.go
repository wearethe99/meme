package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
)

type ServerConfig struct {
	Token string
	Debug bool
	Inline func(*tgbotapi.BotAPI, *tgbotapi.InlineQuery)
	NotFound func(*tgbotapi.BotAPI, tgbotapi.Update)
}

func NewServer(cfg ServerConfig) error {
	bot, err := tgbotapi.NewBotAPI(cfg.Token)

	if err != nil {
		return errors.Wrap(err, "cant start telegram server")
	} else {
		fmt.Printf("Authorized on account <%s>", bot.Self.UserName)
	}

	bot.Debug = cfg.Debug
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.InlineQuery != nil && cfg.Inline != nil {
			cfg.Inline(bot, update.InlineQuery)
		} else {
			cfg.NotFound(bot, update)
		}
	}

	return nil
}
