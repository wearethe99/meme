package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"meme/container"
)

func PrintAssets(_ *cli.Context) error {
	for _, asset := range container.Assets {
		fmt.Println(asset.Prefix + "\t" + asset.Path + "\t" + asset.Image.Hash)
	}

	return nil
}
