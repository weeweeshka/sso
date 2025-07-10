package main

import (
	"github.com/weeweeshka/sso/internal/app"
	"github.com/weeweeshka/sso/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	slogger := SetupLogger()
	slog.Info("Logger started")

	application := app.New(cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL, slogger)
	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	slog.Info("Gracefully shutting down...")

	<-stop

	slog.Info("App stopped")
}

func SetupLogger() *slog.Logger {
	var log *slog.Logger
	log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	return log
}
