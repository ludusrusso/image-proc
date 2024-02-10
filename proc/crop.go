package proc

import "image"

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func crop(img image.Image, r image.Rectangle) image.Image {
	if si, ok := img.(SubImager); ok {
		return si.SubImage(r)
	}

	return img
}
