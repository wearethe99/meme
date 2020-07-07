package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"meme/container"
	"meme/internal"
	"meme/internal/source"
)

func Inline(bot *tgbotapi.BotAPI, inline *tgbotapi.InlineQuery) {
	var results []interface{}

	for _, asset := range container.GetBotAssets() {
		id := inline.ID + ":" + asset.Hash
		uri := container.BotMemeUrl(asset, inline.Query)

		switch asset.Codec {
		case source.GIF:
			results = append(results, gif(asset, id, uri))
		default:
			results = append(results, jpeg(asset, id, uri))
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

func gif(asset *internal.Asset, id string, url string) tgbotapi.InlineQueryResultGIF {
	img := tgbotapi.NewInlineQueryResultGIF(id, url)

	img.ThumbURL = url
	img.Width = asset.Width
	img.Height = asset.Height

	return img
}

func jpeg(asset *internal.Asset, id string, url string) tgbotapi.InlineQueryResultPhoto {
	img := tgbotapi.NewInlineQueryResultPhotoWithThumb(id, url, url)

	img.Width = asset.Width
	img.Height = asset.Height

	return img
}
