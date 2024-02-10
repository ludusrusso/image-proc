package main_test

import (
	"bytes"
	"image/png"
	"log"
	"testing"

	_ "embed"

	"github.com/nfnt/resize"
	"github.com/stretchr/testify/assert"
)

//go:embed test_data/jaga.png
var input []byte

func TestReadImageSize(t *testing.T) {
	// transfor input to io.Reader
	in := bytes.NewReader(input)

	img, err := png.Decode(in)
	if err != nil {
		log.Fatal(err)
	}

	size := img.Bounds().Size()

	assert.Equal(t, 510, size.X)
	assert.Equal(t, 516, size.Y)
}

func TestResizeImage(t *testing.T) {
	// transfor input to io.Reader
	in := bytes.NewReader(input)

	img, err := png.Decode(in)
	if err != nil {
		log.Fatal(err)
	}

	m := resize.Resize(uint(img.Bounds().Size().X)/2, 0, img, resize.Lanczos3)

	assert.Equal(t, 255, m.Bounds().Size().X)
	assert.Equal(t, 258, m.Bounds().Size().Y)
}
