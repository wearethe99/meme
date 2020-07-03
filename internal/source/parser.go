package source

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/jpoz/gomeme"
	"github.com/pkg/errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
)

type codecParser func(io.Reader) (gomeme.Memeable, error)

var codecParserMap = map[Codec]codecParser{
	JPEG: jpegParser,
	GIF: gifParser,
	PNG: pngParser,
}

func detectCodec(input []byte) Codec {
	contentType := detectContentType(input)

	switch contentType {
	case "image/gif":
		return GIF
	case "image/jpeg":
		return JPEG
	case "image/png":
		return PNG
	default:
		return UNKNOWN
	}
}

func detectContentType(input []byte) string {
	return http.DetectContentType(input)
}

func detectHash(input []byte) string {
	hash := md5.New()
	hash.Write(input)
	return hex.EncodeToString(hash.Sum(nil))
}

func jpegParser(reader io.Reader) (gomeme.Memeable, error) {
	if img, err := jpeg.Decode(reader); err != nil {
		return nil, errors.Wrap(err, "cant decode JPEG")
	} else {
		return gomeme.JPEG{Image: img}, nil
	}
}

func gifParser(reader io.Reader) (gomeme.Memeable, error) {
	if img, err := gif.DecodeAll(reader); err != nil {
		return nil, errors.Wrap(err, "cant decode GIF")
	} else {
		return gomeme.GIF{GIF: img}, nil
	}
}

func pngParser(reader io.Reader) (gomeme.Memeable, error) {
	if img, err := png.Decode(reader); err != nil {
		return nil, errors.Wrap(err, "cant decode PNG")
	} else {
		return gomeme.PNG{Image: img}, nil
	}
}