package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/nfnt/resize"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

//go:embed test_data/jaga.png
var input []byte

//go:embed test_data/jaga2.png
var input2 []byte

//go:embed test_data/jaga3.png
var input3 []byte

func resizeImg(w http.ResponseWriter, req *http.Request) {
	// golang atoi

	in := bytes.NewReader(input)
	img, err := png.Decode(in)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ww, err := strconv.Atoi(req.URL.Query().Get("w"))
	if err != nil {
		ww = img.Bounds().Dx()
	}

	hh, err := strconv.Atoi(req.URL.Query().Get("h"))
	if err != nil {
		hh = img.Bounds().Dy()
	}

	slog.Info("sizes", "width", ww, "height", hh)

	m := resize.Resize(uint(ww), 0, img, resize.Lanczos2)
	m.Bounds()

	crpped := m.(SubImager).SubImage(image.Rect(0, 0, ww, hh))

	w.Header().Set("X-size", crpped.Bounds().Size().String())
	w.WriteHeader(http.StatusOK)
	if err := jpeg.Encode(w, crpped, nil); err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", resizeImg)
	http.HandleFunc("/face", faceCrop)

	fmt.Printf("Server started at http://localhost:8080\n")
	http.ListenAndServe(":8080", nil)
}
