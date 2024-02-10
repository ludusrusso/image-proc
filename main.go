package main

import (
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	// open "test.jpg"
	file, err := os.Open("jaga.png")
	if err != nil {
		log.Fatal(err)
	}

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	orgW := img.Bounds().Size().X
	orgH := img.Bounds().Size().Y

	log.Printf("Original Size: %d, %d\n", orgW, orgH)

	m := resize.Resize(uint(orgW)/2, 0, img, resize.Lanczos3)

	out, err := os.Create("jaga_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
