package loader

import (
	"errors"
	"image"
)

type Loader interface {
	Load(path string) (image.Image, error)
	Store(path string, img image.Image) error
}

var ErrNotFound = errors.New("not found")
