package internal

import (
	"github.com/jpoz/gomeme"
	"io"
	"meme/internal/source"
	"meme/utl"
	"regexp"
)


// split the text to the <Top text> and <Bottom> text by
// new line or "//" chars
var textSplitRe = regexp.MustCompile(`\n|/{2}`)

type MemeConfig struct {
	Image *source.Image
	Text string
}

func Meme2Writer(cfg MemeConfig, output io.Writer) error {
	config := gomeme.NewConfig()
	parts := split(cfg.Text)

	switch len(parts) {
	case 1:
		config.BottomText = parts[0]
	case 2:
		config.TopText = parts[0]
		config.BottomText = parts[1]
	}

	config.FontSize = 100

	meme := &gomeme.Meme{
		Config: config,
		Memeable: cfg.Image.Memeable,
	}

	if err := meme.Write(output); err != nil {
		return err
	}

	return nil
}

// returns slice of strings
// * len(2) - ["top text", "bottom text"]
// * len(1) - ["bottom text"]
func split(text string) []string {
	result := textSplitRe.Split(text, 2)

	for ix, value := range result {
		result[ix] = utl.Squish(value)
	}

	return result
}