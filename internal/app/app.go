package app

import (
	grpcapp "github.com/weeweeshka/sso/internal/app/grpc"
	"github.com/weeweeshka/sso/internal/services/auth"
	"github.com/weeweeshka/sso/internal/storage/sqlite"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(grpcPort int, storagePath string, tokenTTL time.Duration, slog *slog.Logger) *App {

	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(slog, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(grpcPort, slog, authService)

	return &App{GRPCServer: grpcApp}
}
