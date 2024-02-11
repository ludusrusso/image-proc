package loader_test

import (
	"image"
	"testing"

	"github.com/ludusrusso/image-proc/pkg/loader"
	"github.com/stretchr/testify/assert"
)

func TestStoreAndLoad(t *testing.T) {
	l := createTestLoader(t)

	img := createTestImage(t)

	err := l.Store("test.png", img)
	if err != nil {
		t.Fatal(err)
	}

	res, err := l.Load("test.png")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, img, res)
}

func TestLoadNotExists(t *testing.T) {
	l := createTestLoader(t)

	_, err := l.Load("not_exists.png")
	assert.ErrorIs(t, err, loader.ErrNotFound)
}

func TestStoreMultipleTimes(t *testing.T) {
	l := createTestLoader(t)

	img := createTestImage(t)

	err := l.Store("test.png", img)
	if err != nil {
		t.Fatal(err)
	}

	err = l.Store("test.png", img)
	if err != nil {
		t.Fatal(err)
	}
}

func createTestLoader(t *testing.T) loader.Loader {
	t.Helper()

	dir := t.TempDir()

	l := loader.NewFileLoader(dir)

	return l
}

func createTestImage(t *testing.T) image.Image {
	t.Helper()

	return image.NewGray(image.Rect(0, 0, 10, 10))
}
