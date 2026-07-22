package repository

import (
	"context"
	"task_tracker/internal/domain/types"
)

type TaskRepository interface {
	Ping(ctx context.Context) error

	Create(ctx context.Context, task *types.Task) error

	List(ctx context.Context, statis *string, limit, offset int) ([]types.Task, error)

	GetByID(ctx context.Context, id uint) (*types.Task, error)
	
	Update(ctx context.Context, id uint, title *string, description *string, status *string) (*types.Task, error)

	Delete(ctx context.Context, id uint) error
}
