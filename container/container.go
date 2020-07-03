package container

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"meme/internal/source"
	"path/filepath"
)

var Sources []*source.Image
var SourceIx = make(map[string]*source.Image)

func LoadSources(dir string) error {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return errors.Wrap(err, "cant load source files")
	}

	for _, f := range files {
		image, err := source.NewImageFromFile(filepath.Join(dir, f.Name()))

		if err != nil {
			return err
		}

		Sources = append(Sources, image)
	}

	return nil
}

func IndexSources() {
	for _, image := range Sources {
		SourceIx[image.Hash] = image
	}
}