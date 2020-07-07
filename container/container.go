package container

import "meme/internal"

var Assets []*internal.Asset
var AssetPrefixGroup = make(map[string][]*internal.Asset)
var AssetHashIx = make(map[string]*internal.Asset)

var BotPrefix string
var BotToken string
var BotText string
