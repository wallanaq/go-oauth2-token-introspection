package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/wallanaq/go-oauth2-token-introspection/internal/server"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		slog.Warn("received shutdown signal", "signal", sig)
		cancel()
	}()

	if err := server.Start(ctx); err != nil {
		slog.Error("server encoutered an error", slog.String("error", err.Error()))
		os.Exit(1)
	}

	slog.Info("server has shut down gracefully")

}
