package database

import (
	"context"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type GormDBInterface interface {
	Create(value interface{}) GormDBInterface
	Find(out interface{}, where ...interface{}) GormDBInterface
	First(dest interface{}, conds ...interface{}) GormDBInterface
	Save(value interface{}) GormDBInterface
	Delete(value interface{}, conds ...interface{}) GormDBInterface
	WithContext(ctx context.Context) GormDBInterface
	GetDB() *gorm.DB
}

type GormDBWrapper struct {
	DB *gorm.DB
	*mock.Mock
}

func (w *GormDBWrapper) Create(value interface{}) GormDBInterface {
	if w.Mock != nil {
		w.Mock.Called(value)
		return w
	}

	w.DB = w.DB.Create(value)
	return w
}

func (w *GormDBWrapper) Find(out interface{}, where ...interface{}) GormDBInterface {
	if w.Mock != nil {
		w.Mock.Called(out, where)
		return w
	}

	w.DB = w.DB.Find(out, where...)
	return w
}

func (w *GormDBWrapper) Delete(value interface{}, conds ...interface{}) GormDBInterface {
	if w.Mock != nil {
		w.Mock.Called(value, conds)
		return w
	}

	w.DB = w.DB.Delete(value, conds)
	return w
}

func (w *GormDBWrapper) First(dest interface{}, conds ...interface{}) GormDBInterface {
	if w.Mock != nil {
		w.Mock.Called(dest, conds)
		return w
	}

	w.DB = w.DB.First(dest, conds)
	return w
}

func (w *GormDBWrapper) Save(value interface{}) GormDBInterface {
	if w.Mock != nil {
		w.Mock.Called(value)
		return w
	}

	w.DB = w.DB.Save(value)
	return w
}

func (w *GormDBWrapper) WithContext(ctx context.Context) GormDBInterface {
	if w.Mock != nil {
		w.Mock.Called(ctx)
		return w
	}

	w.DB = w.DB.WithContext(ctx)
	return w
}

func (w *GormDBWrapper) GetDB() *gorm.DB {
	return w.DB
}
