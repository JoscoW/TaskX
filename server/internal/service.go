package internal

import (
	"context"
	"go.uber.org/fx"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (t *TaskXService) GetTasks(context.Context, *emptypb.Empty) (*pb.GetTasksResp, error) {

	tasks, err := t.Repository.GetTasks()
	if err != nil {
		return nil, err
	}

	return &pb.GetTasksResp{
		Tasks: tasks,
	}, nil
}
