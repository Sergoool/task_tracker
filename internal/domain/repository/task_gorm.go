package repository

import (
	"context"

	"gorm.io/gorm"
)

type TaskGormRepository struct {
	db *gorm.DB
}

func NewTaskGormRepository(db *gorm.DB) *TaskGormRepository {
	return &TaskGormRepository{db: db}
}

func (r *TaskGormRepository) Ping(ctx context.Context) error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.PingContext(ctx)
}
