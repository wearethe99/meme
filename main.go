package main

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"meme/cmd"
	"meme/internal"
	"os"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: internal.SourceDirFlag,
				Value: "assets/lukashenko",
				Usage: "specify a directory with source images",
				EnvVars: []string{"SOURCE_DIR"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "server",
				Aliases: []string{"s"},
				Usage:   "run HTTP server for serving MEME images",
				Before: cmd.LoadAssets,
				Action: cmd.Server,
			},{
				Name:    "bot",
				Aliases: []string{"b"},
				Usage:   "run telegram bot",
				Flags: []cli.Flag {
					&cli.StringFlag{
						Name: internal.BotTokenFlag,
						Usage: "telegram bot token",
						EnvVars: []string{"BOT_TOKEN"},
					},
				},
				Before: cmd.LoadAssets,
				Action: cmd.Bot,
			},{
				Name:    "print",
				Aliases: []string{"p"},
				Usage:   "print assets hash sums",
				Before: cmd.LoadAssets,
				Action: cmd.PrintAssets,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(errors.Wrap(err, "cant run CLI"))
	}
}
