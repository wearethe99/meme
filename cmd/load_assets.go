package cmd

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"meme/container"
	"meme/internal"
)

func LoadAssets(c *cli.Context) error {
	if err := container.LoadAssetsFile(c.String(internal.AssetsJsonFlag)); err != nil {
		return errors.Wrap(err, "cant read assets json file")
	}

	container.IndexAssetHash()
	container.GroupAssetPrefix()

	return nil
}
