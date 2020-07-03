package cmd

import (
	"github.com/urfave/cli/v2"
	"meme/internal"
	"meme/telegram"
)

func Bot(c *cli.Context) error {
	cfg := telegram.ServerConfig{
		Token: c.String(internal.BotTokenFlag),
		Debug: true,
	}

	cfg.Inline = telegram.Inline
	cfg.NotFound = telegram.NotFound

	return telegram.NewServer(cfg)
}
