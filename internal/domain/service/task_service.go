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
		return nil, Validation(map[string]string{
		"user_id": "must be > 0",
		"title":   "required",
	})
	}

	task := &types.Task{
		Title: title,
		Description: description,
		Status: types.New,
	}

	if err := s.repo.Create(ctx, task); err != nil {
		return nil, Internal(err)
	}
	return task, nil
}
func (s *TaskService) List(ctx context.Context, status *string, limit, offset int) ([]types.Task, error) {
	/*if limit < 0 {
		limit = 20
	}*/
	if limit > 100 || offset < 0 || limit < 0 {
	return nil, Validation(map[string]string{
		"limit":  "must be 1..100",
		"offset": "must be >= 0",
	})
}

	
	tasks, err := s.repo.List(ctx, status, limit, offset)
	if err != nil {
		return nil, Internal(err)
	}
	return tasks, nil
}

func (s *TaskService) GetByID(ctx context.Context, id uint) (*types.Task, error) {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			return nil, NotFound(nil)
		}
		return nil, Internal(err)
	}
	return task, nil
}

func (s *TaskService) Update(ctx context.Context, id uint, title *string, description *string, status *string) (*types.Task, error) {
	if id == 0 {
		return nil, Validation(map[string]string{
		"id":  "must be 1..N",
	})
	}
	if title == nil && status == nil {
		return nil, Validation(map[string]string{
		"title":  "should be text",
		"status": "should be text",
	})}

	if title != nil {
		t := strings.TrimSpace(*title)
		if t == "" {
			return nil,  Validation(map[string]string{
			"title":  "should be text",
		})}
		title = &t
	}

	if description != nil {
		d := strings.TrimSpace(*description)
		if d == "" {
			return nil, Validation(map[string]string{
			"description":  "should be text",})
		}
		description = &d
	}

	task, err := s.repo.Update(ctx, id, title, description, status)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			return nil, NotFound(nil)
		}
		return nil, Internal(err)
	}
	return task, nil
}

func (s *TaskService) Delete(ctx context.Context, id uint) error {
	if id == 0 {
		return Validation(map[string]string{
			"id":"should be > 0",
		})
	}
	err := s.repo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, types.ErrNotFound) {
			return NotFound(err)
		}
		return Internal(err)
	}
	return nil
}


