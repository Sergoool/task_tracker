package repository

import (
	"context"
	"errors"

	//"task_tracker/internal/domain/service"
	"task_tracker/internal/domain/types"

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

func (r *TaskGormRepository) Create(ctx context.Context, task *types.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *TaskGormRepository) List(ctx context.Context) ([]types.Task, error) { 
	var tasks []types.Task
	err := r.db.WithContext(ctx).
		Order("id").
		Find(&tasks).Error
	return tasks, err
}

func (r *TaskGormRepository) GetByID(ctx context.Context, id uint) (*types.Task, error) { // получить по id
	var task types.Task
	err := r.db.WithContext(ctx).First(&task, id).Error // SELECT ... WHERE id=?
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, types.ErrNotFound
		}
		return nil, err
	}
	return &task, nil
}