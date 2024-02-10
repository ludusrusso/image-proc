package loader

import (
	"image"
	"path/filepath"
)

type Cache struct {
	l        Loader
	basePath string
}

func NewCache(l Loader, basePath string) Cache {
	return Cache{l: l, basePath: basePath}
}

func (c Cache) Load(path string) (image.Image, error) {
	p := filepath.Join(c.basePath, path)
	return c.l.Load(p)
}

func (c Cache) Store(path string, img image.Image) error {
	p := filepath.Join(c.basePath, path)
	return c.l.Store(p, img)
}
