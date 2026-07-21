package repository

import "context"

type TaskRepository interface {
	Ping(ctx context.Context) error
}
