package ctrl

import (
	"github.com/gin-gonic/gin"
	"meme/container"
	"meme/internal"
	"meme/internal/source"
)

const MemeRoute = "/meme/:hash"

func Meme(c *gin.Context) {
	term := c.Query("term")
	hash := c.Param("hash")
	image, ok := container.SourceIx[hash]

	if !ok {
		c.AbortWithStatusJSON(404, gin.H{
			"error": "meme not found",
		})
	} else {
		memeResponse(c, term, image)
	}
}

func memeResponse(c *gin.Context, term string, image *source.Image) {
	c.Status(200)

	c.Writer.Header().Set("Content-type", image.ContentType)
	c.Writer.Header().Set("Content-Disposition", `inline`)

	cfg := internal.MemeConfig{
		Image: image,
		Text: term,
	}

	_ = internal.Meme2Writer(cfg, c.Writer)
}