package container

import (
	"log/slog"

	"github.com/ludusrusso/image-proc/pkg/cache"
	"github.com/ludusrusso/image-proc/pkg/loader"
)

type Cfg struct {
	LoaderPath string
	CachePath  string
}

type Container struct {
	cfg    Cfg
	cache  *cache.Cache
	loader loader.Loader
	log    *slog.Logger
}

func NewContainer(cfg Cfg) Container {
	return Container{
		cfg: cfg,
	}
}

func (c Container) Logger() *slog.Logger {
	if c.log == nil {
		c.log = slog.New(&slog.TextHandler{})
	}
	return c.log
}

func (c Container) Loader() loader.Loader {
	if c.loader == nil {
		c.loader = loader.NewFileLoader(c.cfg.LoaderPath)
	}
	return c.loader
}

func (c Container) Cache() *cache.Cache {
	if c.cache == nil {
		cl := loader.NewFileLoader(c.cfg.CachePath)
		c.cache = cache.NewCache(cl)
	}
	return c.cache
}
