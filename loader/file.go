package loader

import (
	"image"
	"image/png"
	"log"
	"os"
	"strings"
)

type fileLoader struct{}

func NewFileLoader() Loader {
	return &fileLoader{}
}

func (fl *fileLoader) Load(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return png.Decode(f)
}

func (fl *fileLoader) Store(path string, data image.Image) error {
	// create dir if not exists
	dir := path[:strings.LastIndex(path, "/")]
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	if err := png.Encode(f, data); err != nil {
		return err
	}

	return nil
}
