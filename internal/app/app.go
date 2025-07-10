package app

import (
	grpcapp "github.com/weeweeshka/sso/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(grpcPort int, storagePath string, tokenTTL time.Duration, slog *slog.Logger) *App {
	// TODO: init storage

	// TODO: init auth service

	grpcApp := grpcapp.New(grpcPort, slog)

	return &App{GRPCServer: grpcApp}
}
