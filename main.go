package main

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"meme/cmd"
	"meme/internal"
	"os"
)

type action func (c *cli.Context) error

func main() {
	app := &cli.App{
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: internal.AssetsJsonFlag,
				Value: "assets.json",
				Usage: "specify a path to the assets.json file",
				EnvVars: []string{"ASSETS"},
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
					&cli.StringFlag{
						Name: internal.BotPrefixFlag,
						Usage: "prefix for source images to serve",
						EnvVars: []string{"BOT_PREFIX"},
					},
					&cli.StringFlag{
						Name: internal.BotTextFlag,
						Usage: "start text",
						EnvVars: []string{"BOT_TEXT"},
					},
				},
				Before: cli.BeforeFunc(chain([]action{
					cmd.LoadAssets,
					cmd.ReadBotFlags,
				})),
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
		panic(errors.Wrap(err, "cant run the CLI"))
	}
}

// make a single function which wraps all function from a list
func chain(list []action) action {
	return func(c *cli.Context) error {
		for _, action := range list {
			if err := action(c); err != nil {
				return err
			}
		}

		return nil
	}
}
