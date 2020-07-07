package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"meme/container"
	"meme/internal/source"
	"net/url"
)

func Inline(bot *tgbotapi.BotAPI, inline *tgbotapi.InlineQuery) {
	var results []interface{}

	for ix, image := range container.Sources {
		slug := image.Hash + "?term=" + url.PathEscape(inline.Query)
		id := inline.ID + ":" + string(ix) + ":" + image.Hash
		uri := "https://ivote.live/meme/" + slug

		switch image.Codec {
		case source.GIF:
			results = append(results, gif(image, id, uri))
		default:
			results = append(results, jpeg(image, id, uri))
		}

	}

	cfg := tgbotapi.InlineConfig{
		InlineQueryID: inline.ID,
		IsPersonal: false,
		CacheTime: 300,
		Results: results,
	}

	if _, err := bot.AnswerInlineQuery(cfg); err != nil {
		log.Println(err)
	}
}

func gif(image *source.Image, id string, url string) tgbotapi.InlineQueryResultGIF {
	img := tgbotapi.NewInlineQueryResultGIF(id, url)

	img.ThumbURL = url
	img.Width = image.Width
	img.Height = image.Height

	return img
}

func jpeg(image *source.Image, id string, url string) tgbotapi.InlineQueryResultPhoto {
	img := tgbotapi.NewInlineQueryResultPhotoWithThumb(id, url, url)

	img.Width = image.Width
	img.Height = image.Height

	return img
}