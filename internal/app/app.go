package app

import (
	grpcapp "sso/internal/app/grpc"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: init storage

	// TODO: init auth service

	grpcApp := grpcapp.New(grpcPort)

	return &App{GRPCServer: grpcApp}
}
