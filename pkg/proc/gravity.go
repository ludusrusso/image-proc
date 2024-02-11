package proc

import (
	"image"

	"github.com/ludusrusso/image-proc/pkg/config"
	"github.com/ludusrusso/image-proc/pkg/proc/internal/face"
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
