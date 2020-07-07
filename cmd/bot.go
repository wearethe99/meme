package cmd

import (
	"github.com/urfave/cli/v2"
	"meme/container"
	"meme/telegram"
)

func Bot(_ *cli.Context) error {
	cfg := telegram.ServerConfig{
		Token: container.BotToken,
		Debug: true,
	}

	cfg.Inline = telegram.Inline
	cfg.NotFound = telegram.NotFound

	return telegram.NewServer(cfg)
}
