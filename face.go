package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"

	pigo "github.com/esimov/pigo/core"

	_ "embed"
)

//go:embed facefinder
var cascadeFile []byte

func faceCrop(w http.ResponseWriter, req *http.Request) {
	// golang atoi

	in := bytes.NewReader(input3)
	img, err := png.Decode(in)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	img.Bounds()

	rect := detectFaces(img)

	// pixels := pigo.RgbToGrayscale(img)
	// cols, rows := img.Bounds().Max.X, img.Bounds().Max.Y

	crpped := img.(SubImager).SubImage(rect)

	w.Header().Set("X-size", crpped.Bounds().Size().String())
	w.WriteHeader(http.StatusOK)
	if err := jpeg.Encode(w, crpped, nil); err != nil {
		return
	}
}

func detectFaces(img image.Image) image.Rectangle {
	pixels := pigo.RgbToGrayscale(img)
	cols, rows := img.Bounds().Max.X, img.Bounds().Max.Y

	pp := pigo.NewPigo()
	// Unpack the binary file. This will return the number of cascade trees,
	// the tree depth, the threshold and the prediction from tree's leaf nodes.
	classifier, err := pp.Unpack(cascadeFile)
	if err != nil {
		log.Fatalf("Error reading the cascade file: %s", err)
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

	// Run the classifier over the obtained leaf nodes and return the detection results.
	// The result contains quadruplets representing the row, column, scale and detection score.
	dets := classifier.RunCascade(cParams, angle)

	// Calculate the intersection over union (IoU) of two clusters.
	dets = classifier.ClusterDetections(dets, 0.2)

	maxScore := dets[0]
	for _, det := range dets {
		if det.Q > maxScore.Q {
			maxScore = det
		}
	}

	r := maxScore.Scale/2 + 100
	yc := maxScore.Row
	xc := maxScore.Col
	return image.Rect(xc-r, yc-r, xc+r, yc+r)
}
