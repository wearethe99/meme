package container

import (
	"meme/internal"
	"net/url"
)

func GetBotAssets() []*internal.Asset {
	if ary, ok := AssetPrefixGroup[BotPrefix]; !ok {
		return nil
	} else {
		return ary
	}
}

func BotMemeUrl(asset *internal.Asset, term string) string {
	return internal.Website + "/memes/" + BotPrefix + "/" + asset.Hash + "?term=" + url.QueryEscape(term)
}
