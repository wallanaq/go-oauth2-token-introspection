package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

func Start(ctx context.Context) error {

	srv := &http.Server{
		Addr: ":8080",
	}

	serverErr := make(chan error, 1)

	go func() {
		slog.Info("starting http server", slog.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		return err
	case <-ctx.Done():
		slog.Info("shutting down server")
		return shutdown(srv)
	}

}

func shutdown(srv *http.Server) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("error shutting down server", slog.String("error", err.Error()))
		return err
	}

	return nil

}
