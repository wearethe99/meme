package cmd

import (
	"github.com/urfave/cli/v2"
	"meme/container"
	"meme/internal"
)

func LoadAssets(c *cli.Context) error {
	err := container.LoadSources(c.String(internal.SourceDirFlag))
	container.IndexSources()

	if err != nil {
		return err
	}

	return nil
}
