package service

import (
	"context"

	"task_tracker/internal/domain/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}


func (s *TaskService) Health(ctx context.Context) error {
	return s.repo.Ping(ctx)
}

func (s *TaskService) Version() string {
	return "1.0"
}
