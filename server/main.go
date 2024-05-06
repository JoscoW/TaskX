package server

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"taskx-server/internal"
	"taskx-server/pb"
	"taskx-server/repo"
)

func main() {
	fx.New(
		fx.Provide(
			repo.NewConfig,
			repo.NewTaskXRepository,
			fx.Annotate(
				internal.NewGRPCServer,
				fx.As(new(grpc.ServiceRegistrar)),
			),
			fx.Annotate(
				internal.NewTaskXService,
				fx.As(new(pb.TaskXServiceServer)),
			),
		),
		fx.Invoke(
			pb.RegisterTaskXServiceServer,
		),
	).Run()
}
