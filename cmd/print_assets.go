package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"meme/container"
)

func PrintAssets(c *cli.Context) error {
	for _, image := range container.Sources {
		fmt.Println(image.Hash + "\t" + image.Name)
	}

	return nil
}
