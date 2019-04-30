package main

import (
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	bpActivity "github.com/yumuranaoki/microservice-with-grpc/proto/activity"
	bpProject "github.com/yumuranaoki/microservice-with-grpc/proto/project"
	bpTask "github.com/yumuranaoki/microservice-with-grpc/proto/task"
)

type TaskService struct {
	store          Store
	activityClient pbActivity.ActivityServiceClient
	projectClient  pbProject.ProjectServiceClient
}

func (s *TaskService) CreateTask(
	ctx context.Context,
	req *pbTask.CreateTaskRequest
) (*pbTask.CreateTaskResponse, error) {
	if reg.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "empty task name")
	}
	resp, err := s.projectClient.FindProject(ctx, &pbProject.FindProject{ProjectId: req.GetProjectId()})
	if err != nil {
		return
			nil,
			status.Error(
				codes.NotFound, err.Error()
			)
	}
	userID := md.GetUserIDFromContext(ctx)
	now := ptypes.TimestampNow()

	task, err := s.store.CreateTask(&pbTask.Task{
		Name: req.GetName(),
		Status: pbTask.Status_WAITING,
		UserId: userID,
		ProjectId: resp.ProjectId.GetId(),
		CreatedAt: now,
		UpdatedAt: now,
	})
}