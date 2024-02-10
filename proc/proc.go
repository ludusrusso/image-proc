package proc

import (
	"image"

	"github.com/ludusrusso/image-proc/config"
)

func ProcImage(cnf config.ProcConfig, img image.Image) image.Image {
	grav := computeGravity(img, cnf.Gravity)

	if cnf.Crop == config.CropTypeCrop {
		return crop(img, grav)
	}

	if cnf.Crop == config.CropTypeScale {
		return resize(img, cnf.Width, cnf.Height)
	}

	if cnf.Crop == config.CropTypeThumb {

		img = crop(img, grav)
		return resize(img, cnf.Width, cnf.Height)
	}

	return img
}
