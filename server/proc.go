package server

import (
	"image"
	"image/jpeg"
	"log/slog"
	"net/http"
	"strings"

	"github.com/ludusrusso/image-proc/pkg/cache"
	"github.com/ludusrusso/image-proc/pkg/config"
	"github.com/ludusrusso/image-proc/pkg/loader"
	"github.com/ludusrusso/image-proc/pkg/proc"
)

type ImageProc struct {
	cache  *cache.Cache
	loader loader.Loader
	log    *slog.Logger
}

func (i ImageProc) handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		cnf, path := getConfigAndPath(req)

		img, err := i.loadAndProcImage(cnf, path)
		if err != nil {
			slog.Error("load and proc image", "error", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := jpeg.Encode(w, img, nil); err != nil {
			return
		}
	})
}

func (i ImageProc) loadAndProcImage(cnf config.Config, path string) (image.Image, error) {
	img, ok := i.cache.Search(cnf, path)
	if ok {
		slog.Info("cache hit", "path", path, "config", cnf.String())
		return img, nil
	}

	img, err := i.loader.Load(path)
	if err != nil {
		return nil, err
	}

	res := proc.ProcImage(cnf, img)
	i.cache.Store(cnf, path, res)

	return res, nil
}

func getConfigAndPath(req *http.Request) (config.Config, string) {
	path := req.URL.Path
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	res := strings.SplitN(path, "/image/", 2)
	if len(res) != 2 {
		return config.Config{}, ""
	}

	return config.Parse(res[0]), res[1]
}
