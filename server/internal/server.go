package internal

import (
	"context"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

func NewGRPCServer(lc fx.Lifecycle) *grpc.Server {
	srv := grpc.NewServer()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			slog.Info("Starting gRPC server")

			ln, err := net.Listen("tcp", ":9000")
			if err != nil {
				return err
			}

			go func() {
				if err := srv.Serve(ln); err != nil {
					slog.Error("Failed to Serve gRPC", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			slog.Info("Gracefully stopping gRPC server")

			srv.GracefulStop()

			return nil
		},
	})

	return srv
}
