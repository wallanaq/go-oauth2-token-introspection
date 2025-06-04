package main

import (
	"log/slog"
	"os"

	"github.com/wallanaq/go-oauth2-token-introspection/internal/app"
)

func main() {

	if err := app.Run(); err != nil {
		slog.Error("error running app", slog.String("error", err.Error()))
		os.Exit(1)
	}

}
