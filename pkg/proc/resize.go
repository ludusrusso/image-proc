package proc

import (
	"image"

	r "github.com/nfnt/resize"
)

func resize(img image.Image, w, h uint) image.Image {
	return r.Resize(w, h, img, r.Lanczos2)
}
