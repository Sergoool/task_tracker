package service

import (
	"context"
	"errors"
	"strings"

	"task_tracker/internal/domain/repository"
	"task_tracker/internal/domain/types"
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

func (s *TaskService) Create(ctx context.Context, title string,
description string) (*types.Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, types.ErrValidation
	}

	task := &types.Task{
		Title: title,
		Description: description,
		Status: types.New,
	}

	if err := s.repo.Create(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}
func (s *TaskService) List(ctx context.Context) ([]types.Task, error) {
	return s.repo.List(ctx)
}

func (s *TaskService) GetByID(ctx context.Context, id uint) (*types.Task, error) {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			return nil, types.ErrNotFound
		}
		return nil, err
	}
	return task, nil
}
