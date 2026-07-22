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

func (r *TaskGormRepository) List(ctx context.Context, status *string, limit, offset int) ([]types.Task, error) {
	var tasks []types.Task

	q := r.db.WithContext(ctx).Model(&types.Task{}).Order("id")

	if status != nil {
		q = q.Where("status = ?", *status) // where запрос
	}
	if limit > 0 {
		q = q.Limit(limit) //limit
	}
	if offset > 0 {
		q = q.Offset(offset) //offset
	}

	err := q.Find(&tasks).Error
	return tasks, err
}

func (r *TaskGormRepository) Update(ctx context.Context, id uint, title *string, status *string) (*types.Task, error) {
	task, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if title != nil {
		task.Title = *title
	}
	if status != nil {
		task.Status = *status
	}

	if err := r.db.WithContext(ctx).Save(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskGormRepository) Delete(ctx context.Context, id uint) error {
	res := r.db.WithContext(ctx).Delete(&types.Task{}, id) //delete where
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return types.ErrNotFound
	}
	return nil
}