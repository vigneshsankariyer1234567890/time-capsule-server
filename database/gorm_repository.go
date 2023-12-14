package database

import (
	"context"
)

// Generic database operations interface
type Repository[T any] interface {
	Create(ctx context.Context, obj T) error
	Find(ctx context.Context, conds ...interface{}) ([]T, error)
	First(ctx context.Context, conds ...interface{}) (T, error)
	Save(ctx context.Context, obj T) error
	Delete(ctx context.Context, obj T) error
}

// GormDatabase is a GORM implementation of the Repository interface
type GormRepository[T any] struct {
	DB *GormDBWrapper
}

// NewGormDatabase creates a new instance of GormDatabase
func NewGormDatabase[T any](db *GormDBWrapper) *GormRepository[T] {
	return &GormRepository[T]{DB: db}
}

func (gdb *GormRepository[T]) Create(ctx context.Context, obj T) error {
	var with_context = gdb.DB.WithContext(ctx)
	var create = with_context.Create(&obj)
	var db = create.GetDB()
	return db.Error
}

func (gdb *GormRepository[T]) Find(ctx context.Context, conds ...interface{}) ([]T, error) {
	var results []T
	err := gdb.DB.WithContext(ctx).Find(&results, conds...).GetDB().Error
	return results, err
}

func (gdb *GormRepository[T]) First(ctx context.Context, conds ...interface{}) (T, error) {
	var result T
	err := gdb.DB.WithContext(ctx).First(&result, conds...).GetDB().Error
	return result, err
}

func (gdb *GormRepository[T]) Save(ctx context.Context, obj T) error {
	return gdb.DB.WithContext(ctx).Save(&obj).GetDB().Error
}

func (gdb *GormRepository[T]) Delete(ctx context.Context, obj T) error {
	return gdb.DB.WithContext(ctx).Delete(&obj).GetDB().Error
}
