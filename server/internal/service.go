package internal

import (
	"go.uber.org/fx"
	"taskx-server/pb"
	"taskx-server/repo"
)

func NewTaskXService(_ fx.Lifecycle, da *repo.TaskXRepository) *TaskXService {
	return &TaskXService{
		Repository: da,
	}
}

type TaskXService struct {
	pb.UnimplementedTaskXServiceServer
	Repository *repo.TaskXRepository
}
