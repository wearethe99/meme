package ctrl

import (
	"github.com/gin-gonic/gin"
	"meme/container"
	"meme/internal"
)

const MemeRoute = "/memes/:prefix/:hash"

func Meme(c *gin.Context) {
	term, prefix, hash := c.Query("term"), c.Param("prefix"), c.Param("hash")
	asset, ok := container.FindAssetByPrefixAndHash(prefix, hash)

	if !ok {
		memeNotFoundResponse(c)
	} else {
		memeResponse(c, term, asset)
	}
}

func memeNotFoundResponse(c *gin.Context) {
	c.AbortWithStatusJSON(404, gin.H{
		"error": "meme not found",
	})
}

func memeResponse(c *gin.Context, term string, asset *internal.Asset) {
	c.Status(200)

	c.Writer.Header().Set("Content-type", asset.Image.ContentType)
	c.Writer.Header().Set("Content-Disposition", `inline`)

	cfg := internal.MemeConfig{
		Image: asset.Image,
		Text: term,
	}

	_ = internal.Meme2Writer(cfg, c.Writer)
}