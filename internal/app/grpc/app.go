package grpcapp

import (
	"fmt"
	authgrpc "github.com/weeweeshka/sso/internal/grpc/auth"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	slog       *slog.Logger
	port       int
}

func New(port int, slog *slog.Logger, authService authgrpc.Auth) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.RegisterServer(gRPCServer, authService)

	return &App{
		gRPCServer: gRPCServer,
		slog:       slog,
		port:       port,
	}
}

func (a *App) Run() error {
	const op = "grpcsapp.Run"
	log := a.slog.With(slog.String("op", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	log.Info("gRPC server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	const op = "grpcsapp.Stop"

	a.slog.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))
	a.gRPCServer.GracefulStop()
}
