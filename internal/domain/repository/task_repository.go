package repository

import (
	"context"
	"task_tracker/internal/domain/types"
)

type TaskRepository interface {
	Ping(ctx context.Context) error

	Create(ctx context.Context, task *types.Task) error

	List(ctx context.Context) ([]types.Task, error)

	GetByID(ctx context.Context, id uint) (*types.Task, error)
	
}
