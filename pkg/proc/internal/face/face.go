package face

import (
	"image"

	pigo "github.com/esimov/pigo/core"

	_ "embed"
)

//go:embed facefinder
var cascadeFile []byte

func Detect(img image.Image) image.Rectangle {
	pixels := pigo.RgbToGrayscale(img)
	cols, rows := img.Bounds().Max.X, img.Bounds().Max.Y

	pp := pigo.NewPigo()

	classifier, err := pp.Unpack(cascadeFile)
	if err != nil {
		return img.Bounds()
	}

	angle := 0.0 // cascade rotation angle. 0.0 is 0 radians and 1.0 is 2*pi radians

	cParams := pigo.CascadeParams{
		MinSize:     20,
		MaxSize:     1000,
		ShiftFactor: 0.1,
		ScaleFactor: 1.1,
		ImageParams: pigo.ImageParams{
			Pixels: pixels,
			Rows:   rows,
			Cols:   cols,
			Dim:    cols,
		},
	}

	dets := classifier.RunCascade(cParams, angle)
	dets = classifier.ClusterDetections(dets, 0.2)

	detection, ok := findBestMatch(dets)
	if !ok {
		return img.Bounds()
	}

	return computeRect(detection)
}

func computeRect(det pigo.Detection) image.Rectangle {
	r := det.Scale/2 + 100
	yc := det.Row
	xc := det.Col
	return image.Rect(xc-r, yc-r, xc+r, yc+r)
}

func findBestMatch(dets []pigo.Detection) (pigo.Detection, bool) {
	if len(dets) == 0 {
		return pigo.Detection{}, false
	}

	maxScore := dets[0]
	for _, det := range dets {
		if det.Q > maxScore.Q {
			maxScore = det
		}
	}

	return maxScore, true
}
