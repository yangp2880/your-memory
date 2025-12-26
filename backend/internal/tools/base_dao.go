package pkg

import (
	"context"

	"gorm.io/gorm"
)

type CRUD[T any] interface {
	Create(ctx context.Context, model *T) error
	Update(ctx context.Context, model *T) error
	Delete(ctx context.Context, id uint64) error
	FindByID(ctx context.Context, id uint64) (*T, error)
	List(ctx context.Context) ([]T, error)
}
type BaseDao[T any] struct {
	db *gorm.DB
}

func NewBaseRepo[T any](db *gorm.DB) *BaseDao[T] {
	return &BaseDao[T]{db: db}
}

func (r *BaseDao[T]) Create(ctx context.Context, model *T) error {
	return nil
}

func (r *BaseDao[T]) Update(ctx context.Context, model *T) error {
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *BaseDao[T]) Delete(ctx context.Context, id uint64) error {
	var t T
	return r.db.WithContext(ctx).Delete(&t, id).Error
}

func (r *BaseDao[T]) FindByID(ctx context.Context, id uint64) (*T, error) {
	var t T
	err := r.db.WithContext(ctx).First(&t, id).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *BaseDao[T]) List(ctx context.Context) ([]T, error) {
	var list []T
	err := r.db.WithContext(ctx).Find(&list).Error
	return list, err
}
