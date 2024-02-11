package server

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/ludusrusso/image-proc/x/container"
)

type Config struct {
	Host string
	Port string
}

func NewServer(
	config Config,
	c container.Container,
) http.Handler {
	mux := http.NewServeMux()

	addRoutes(mux, c)
	var handler http.Handler = mux

	return handler
}

func addRoutes(mux *http.ServeMux, c container.Container) {
	imgPro := ImageProc{
		cache:  c.Cache(),
		loader: c.Loader(),
		log:    c.Logger(),
	}

	mux.Handle("/", imgPro.handler())
}

func RunServer(ctx context.Context, c container.Container, cfg Config) error {
	srv := NewServer(cfg, c)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(cfg.Host, cfg.Port),
		Handler: srv,
	}

	go func() {
		log.Printf("listening on http://%s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		if err := httpServer.Shutdown(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}
