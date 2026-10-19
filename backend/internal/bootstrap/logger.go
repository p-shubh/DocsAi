package bootstrap

import (
	"log/slog"
	"os"
)

func (a *Application) loadLogger() error {
	a.Logger = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	)

	a.Logger.Info("Logger initialized")

	return nil
}