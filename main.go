package main

import (
	_ "embed"
	"fmt"
	"image"
	"image/jpeg"
	"log/slog"
	"net/http"
	"strings"

	"github.com/ludusrusso/image-proc/config"
	"github.com/ludusrusso/image-proc/loader"
	"github.com/ludusrusso/image-proc/proc"
)

func procImage(l loader.Loader, c loader.Cache) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		cnf, path := getConfigAndPath(req)

		img, err := loadAndProcImage(l, c, cnf, path)
		if err != nil {
			slog.Error("load and proc image", "error", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := jpeg.Encode(w, img, nil); err != nil {
			return
		}
	}
}

func main() {
	l := loader.NewFileLoader()
	cache := loader.NewCache(l, "/tmp/cache")

	http.HandleFunc("/", procImage(l, cache))

	fmt.Printf("Server started at http://localhost:8080\n")
	http.ListenAndServe(":8080", nil)
}

func getConfigAndPath(req *http.Request) (config.ProcConfig, string) {
	path := req.URL.Path
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	res := strings.SplitN(path, "/image/", 2)
	if len(res) != 2 {
		return config.ProcConfig{}, ""
	}

	return config.Parse(res[0]), res[1]
}

func loadAndProcImage(l loader.Loader, c loader.Cache, cnf config.ProcConfig, path string) (image.Image, error) {
	img, ok := searchCache(c, cnf, path)
	if ok {
		slog.Info("cache hit", "path", path, "config", cnf.String())
		return img, nil
	}

	img, err := l.Load(path)
	if err != nil {
		return nil, err
	}

	res := proc.ProcImage(cnf, img)
	storeCache(c, cnf, path, res)

	return res, nil
}

func searchCache(c loader.Cache, cnf config.ProcConfig, path string) (image.Image, bool) {
	cp := cachePath(cnf, path)
	img, err := c.Load(cp)
	if err != nil {
		return nil, false
	}
	return img, true
}

func storeCache(c loader.Cache, cnf config.ProcConfig, path string, img image.Image) {
	cp := cachePath(cnf, path)
	if err := c.Store(cp, img); err != nil {
		slog.Error("store cache", "error", err)
	}
}

func cachePath(cnf config.ProcConfig, path string) string {
	return fmt.Sprintf("%s/%s", cnf.String(), path)
}
