package cache

import (
	"fmt"
	"image"
	"log/slog"

	"github.com/ludusrusso/image-proc/pkg/config"
	"github.com/ludusrusso/image-proc/pkg/loader"
)

type Cache struct {
	l loader.Loader
}

func NewCache(l loader.Loader) *Cache {
	return &Cache{
		l: l,
	}
}

func (c *Cache) Search(cnf config.Config, path string) (image.Image, bool) {
	cp := c.path(cnf, path)
	img, err := c.l.Load(cp)
	if err != nil {
		return nil, false
	}
	return img, true
}

func (c *Cache) Store(cnf config.Config, path string, img image.Image) {
	cp := c.path(cnf, path)
	if err := c.l.Store(cp, img); err != nil {
		slog.Error("store cache", "error", err)
	}
}

func (c *Cache) path(cnf config.Config, path string) string {
	return fmt.Sprintf("%s/%s", cnf.String(), path)
}
