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

func (t *TaskXService) AddTask(_ context.Context, req *pb.AddTaskReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, t.Repository.AddTask(req.Description)

}

func (t *TaskXService) CompleteTask(_ context.Context, req *pb.CompleteTaskReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, t.Repository.CompleteTask(req.ID)
}
