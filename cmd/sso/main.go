package sso

import (
	"log/slog"
	"os"
	"sso/internal/app"
	"sso/internal/config"
)

func main() {
	cfg := config.MustLoad()
	SetupLogger()
	slog.Info("Logger started")
	application := app.New(cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	application.GRPCServer.MustRun()
}

func SetupLogger() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
