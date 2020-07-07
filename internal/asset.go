package internal

import "meme/internal/source"

type Asset struct {
	Prefix string
	Path string
	Width int
	Height int
	*source.Image
}

func (a *Asset) Read() error {
	var err error

	a.Image, err = source.NewImageFromFile(a.Path)

	return err
}