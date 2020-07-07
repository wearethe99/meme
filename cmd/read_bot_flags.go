package cmd

import (
	"errors"
	"github.com/urfave/cli/v2"
	"meme/container"
	"meme/internal"
)

func ReadBotFlags(c *cli.Context) error {
	container.BotPrefix = c.String(internal.BotPrefixFlag)
	container.BotToken = c.String(internal.BotTokenFlag)
	container.BotText = c.String(internal.BotTextFlag)

	if container.BotToken == "" {
		return errors.New("bot token is required")
	}

	if container.BotText == "" {
		return errors.New("bot text is required")
	}

	if _, ok := container.AssetPrefixGroup[container.BotPrefix]; !ok {
		return errors.New("specified prefix not exists in the assets")
	}

	return nil
}
