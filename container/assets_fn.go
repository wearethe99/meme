package container

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"meme/internal"
)

func LoadAssetsFile(pathname string) error {
	var assets []*internal.Asset
	file, err := ioutil.ReadFile(pathname)

	if err != nil {
		return errors.Wrap(err, "cant open a file")
	}

	if err = json.Unmarshal(file, &assets); err != nil {
		return errors.Wrap(err, "cant parse a json")
	}

	for _, asset := range assets {
		if err := asset.Read(); err != nil {
			return errors.Wrap(err, "cant read asset image")
		}
	}

	Assets = assets

	return nil
}

func GroupAssetPrefix() {
	data := make(map[string][]*internal.Asset)

	for _, asset := range Assets {
		if ary, ok := data[asset.Prefix]; !ok {
			data[asset.Prefix] = make([]*internal.Asset, 0)
		} else {
			data[asset.Prefix] = append(ary, asset)
		}
	}

	AssetPrefixGroup = data
}

func IndexAssetHash() {
	for _, asset := range Assets {
		AssetHashIx[asset.Image.Hash] = asset
	}
}

func FindAssetByPrefixAndHash(prefix string, hash string) (*internal.Asset, bool) {
	asset, ok := AssetHashIx[hash]

	if !ok {
		return nil, false
	} else if asset.Prefix != prefix {
		return nil, false
	}

	return asset, true
}
