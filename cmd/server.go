package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"meme/ctrl"
)

func Server(c *cli.Context) error {
	api := gin.Default()
	api.GET(ctrl.MemeRoute, ctrl.Meme)

	if err := api.Run(); err != nil {
		return errors.Wrap(err, "cant launch HTTP server")
	}

	return nil
}
