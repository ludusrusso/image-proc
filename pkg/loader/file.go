package loader

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"log/slog"
)

type fileLoader struct {
	baseDir string
}

func NewFileLoader(baseDir string) Loader {
	slog.Info("creating file loader", "baseDir", baseDir)
	return &fileLoader{
		baseDir: baseDir,
	}
}

func (fl *fileLoader) Load(path string) (image.Image, error) {
	path = filepath.Join(fl.baseDir, path)

	slog.Info("loading image", "path", path)

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	defer f.Close()

	return png.Decode(f)
}

func (fl *fileLoader) Store(path string, data image.Image) error {
	path = filepath.Join(fl.baseDir, path)

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
