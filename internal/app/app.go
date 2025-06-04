package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/wallanaq/go-oauth2-token-introspection/internal/config"
	"github.com/wallanaq/go-oauth2-token-introspection/internal/server"
)

func Run() error {

	cfg := config.Load()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		slog.Warn("received shutdown signal", "signal", sig)
		cancel()
	}()

	if err := server.Start(ctx, cfg); err != nil {
		slog.Error("server encoutered an error", slog.String("error", err.Error()))
		return err
	}

	slog.Info("server has shut down gracefully")

	return nil

}
