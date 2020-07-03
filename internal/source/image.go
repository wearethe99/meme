package source

import (
	"bytes"
	"github.com/jpoz/gomeme"
	"github.com/pkg/errors"
	"io/ioutil"
)

type Image struct {
	ContentType string
	Codec 		Codec
	Hash 		string
	Name 		string
	Memeable 	gomeme.Memeable
	Width		int
	Height		int
}

func NewImageFromFile(pathname string) (*Image, error) {
	input, err := ioutil.ReadFile(pathname)

	if err != nil {
		return nil, errors.Wrap(err, "cant read source file")
	}

	codec, contentType := detectCodec(input), detectContentType(input)
	hash := detectHash(input)

	parser, ok := codecParserMap[codec]

	if !ok {
		return nil, errors.Wrap(nil, "cant find a parser for that codec")
	}

	memeable, err := parser(bytes.NewBuffer(input))

	if err != nil {
		return nil, errors.Wrap(err, "decoding error")
	}

	return &Image{
		Name: pathname,
		ContentType: contentType,
		Codec: codec,
		Hash: hash,
		Memeable: memeable,
		Width: 800,
		Height: 500,
	}, nil
}