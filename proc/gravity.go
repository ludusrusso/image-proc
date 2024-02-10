package proc

import (
	"image"

	"github.com/ludusrusso/image-proc/config"
	"github.com/ludusrusso/image-proc/proc/internal/face"
)

func computeGravity(img image.Image, gconf config.GravityType) image.Rectangle {
	if gconf == config.GravityTypeCenter {
		return img.Bounds()
	}
	if gconf == config.GravityTypeFace {
		return face.Detect(img)
	}

	return img.Bounds()
}
