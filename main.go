package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/signal"

	"github.com/ludusrusso/image-proc/server"
	"github.com/ludusrusso/image-proc/x/container"
)

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	c := container.NewContainer(container.Cfg{
		LoaderPath: "./tmp",
		CachePath:  "/tmp/cache",
	})

	return server.RunServer(ctx, c, server.Config{
		Host: "localhost",
		Port: "8080",
	})
}
