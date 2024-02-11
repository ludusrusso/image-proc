package proc

import (
	"bytes"
	"image/png"
	"log"
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed test_data/jaga.png
var input []byte

func TestResizeImage(t *testing.T) {
	in := bytes.NewReader(input)

	img, err := png.Decode(in)
	if err != nil {
		log.Fatal(err)
	}

	m := resize(img, uint(img.Bounds().Size().X)/2, 0)

	assert.Equal(t, 255, m.Bounds().Size().X)
	assert.Equal(t, 258, m.Bounds().Size().Y)
}
