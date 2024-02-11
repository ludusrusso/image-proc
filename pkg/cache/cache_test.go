package cache_test

import (
	"image"
	"testing"

	"github.com/ludusrusso/image-proc/pkg/cache"
	"github.com/ludusrusso/image-proc/pkg/config"
	"github.com/ludusrusso/image-proc/pkg/loader"
	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	c := createCache(t)

	t.Run("StoreAndLoad", func(t *testing.T) {
		img := createTestImage(t)

		c.Store(config.Config{}, "test.png", img)

		res, ok := c.Search(config.Config{}, "test.png")
		assert.True(t, ok)

		assert.Equal(t, img, res)
	})

	t.Run("LoadNotExists", func(t *testing.T) {
		_, ok := c.Search(config.Config{}, "not_found.png")
		assert.False(t, ok)
	})
}

func createCache(t *testing.T) *cache.Cache {
	t.Helper()

	dir := t.TempDir()
	c := loader.NewFileLoader(dir)
	return cache.NewCache(c)
}

func createTestImage(t *testing.T) image.Image {
	t.Helper()

	return image.NewGray(image.Rect(0, 0, 100, 100))
}
