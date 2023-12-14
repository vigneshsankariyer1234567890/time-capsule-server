package database

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TestModel struct {
	ID   int
	Name string
}

var (
	db   *gorm.DB
	repo *GormRepository[TestModel]
)

func TestMain(m *testing.M) {
	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to in-memory database")
	}

	err = db.AutoMigrate(&TestModel{})
	if err != nil {
		panic("failed to auto-migrate")
	}

	wrapper := &GormDBWrapper{DB: db, Mock: new(mock.Mock)}
	repo = NewGormDatabase[TestModel](wrapper)

	// Run the tests
	code := m.Run()

	// Any cleanup if necessary

	os.Exit(code)
}

func resetMock() {
	repo.DB.Mock = new(mock.Mock)
	repo.DB.DB.Error = nil
}

func TestGormRepository_Create_Positive(t *testing.T) {
	resetMock()

	repo.DB.Mock.On("Create", mock.Anything).Return(repo.DB)
	repo.DB.Mock.On("WithContext", mock.Anything).Return(repo.DB)

	ctx := context.Background()
	model := TestModel{Name: "Test Name"}

	err := repo.Create(ctx, model)
	assert.NoError(t, err)
	repo.DB.Mock.AssertExpectations(t)
}

func TestGormRepository_Create_Negative(t *testing.T) {
	resetMock()

	fakeError := errors.New("create error")
	repo.DB.Mock.On("Create", mock.Anything).Run(func(args mock.Arguments) {
		repo.DB.DB.Error = fakeError
	}).Return(repo.DB)
	repo.DB.Mock.On("WithContext", mock.Anything).Return(repo.DB)

	ctx := context.Background()
	model := TestModel{Name: "Test Name"}

	err := repo.Create(ctx, model)
	assert.Error(t, err)
	assert.Equal(t, fakeError, err)
	repo.DB.Mock.AssertExpectations(t)
}

func TestGormRepository_Find_Positive(t *testing.T) {
	resetMock()

	repo.DB.Mock.On("Find", mock.Anything, mock.Anything).Return(repo.DB)
	repo.DB.Mock.On("WithContext", mock.Anything).Return(repo.DB)

	ctx := context.Background()

	results, err := repo.Find(ctx, "Name = ?", "Test Name")
	assert.NoError(t, err)
	assert.Empty(t, results)
	repo.DB.Mock.AssertExpectations(t)
}

func TestGormRepository_Find_Negative(t *testing.T) {
	resetMock()

	fakeError := errors.New("find error")
	repo.DB.Mock.On("Find", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		repo.DB.DB.Error = fakeError
	}).Return(repo.DB)
	repo.DB.Mock.On("WithContext", mock.Anything).Return(repo.DB)

	ctx := context.Background()

	_, err := repo.Find(ctx, "ID = ?", -1)
	assert.Error(t, err)
	assert.Equal(t, fakeError, err)
	repo.DB.Mock.AssertExpectations(t)
}

func TestGormRepository_First_Positive(t *testing.T) {
	resetMock()

	repo.DB.Mock.On("First", mock.Anything, mock.Anything).Return(repo.DB)
	repo.DB.Mock.On("WithContext", mock.Anything).Return(repo.DB)

	ctx := context.Background()

	result, err := repo.First(ctx, "Name = ?", "Test Name")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	repo.DB.Mock.AssertExpectations(t)
}

func TestGormRepository_First_Negative(t *testing.T) {
	resetMock()

	fakeError := errors.New("first error")
	repo.DB.Mock.On("First", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		repo.DB.DB.Error = fakeError
	}).Return(repo.DB)
	repo.DB.Mock.On("WithContext", mock.Anything).Return(repo.DB)

	ctx := context.Background()

	_, err := repo.First(ctx, "ID = ?", -1)
	assert.Error(t, err)
	assert.Equal(t, fakeError, err)
	repo.DB.Mock.AssertExpectations(t)
}

func TestGormRepository_Save_Positive(t *testing.T) {
	resetMock()

	repo.DB.Mock.On("Save", mock.Anything).Return(repo.DB)
	repo.DB.Mock.On("WithContext", mock.Anything).Return(repo.DB)

	ctx := context.Background()
	model := TestModel{ID: 1, Name: "Updated Name"}

	err := repo.Save(ctx, model)
	assert.NoError(t, err)
	repo.DB.Mock.AssertExpectations(t)
}

func TestGormRepository_Save_Negative(t *testing.T) {
	resetMock()

	fakeError := errors.New("save error")
	repo.DB.Mock.On("Save", mock.Anything).Run(func(args mock.Arguments) {
		repo.DB.DB.Error = fakeError
	}).Return(repo.DB)
	repo.DB.Mock.On("WithContext", mock.Anything).Return(repo.DB)

	ctx := context.Background()
	model := TestModel{ID: 1, Name: "Updated Name"}

	err := repo.Save(ctx, model)
	assert.Error(t, err)
	assert.Equal(t, fakeError, err)
	repo.DB.Mock.AssertExpectations(t)
}

func TestGormRepository_Delete_Positive(t *testing.T) {
	resetMock()

	repo.DB.Mock.On("Delete", mock.Anything, mock.Anything).Return(repo.DB)
	repo.DB.Mock.On("WithContext", mock.Anything).Return(repo.DB)

	ctx := context.Background()
	model := TestModel{ID: 1}

	err := repo.Delete(ctx, model)
	assert.NoError(t, err)
	repo.DB.Mock.AssertExpectations(t)
}

func TestGormRepository_Delete_Negative(t *testing.T) {
	resetMock()

	fakeError := errors.New("delete error")
	repo.DB.Mock.On("Delete", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		repo.DB.DB.Error = fakeError
	}).Return(repo.DB)
	repo.DB.Mock.On("WithContext", mock.Anything).Return(repo.DB)

	ctx := context.Background()
	model := TestModel{ID: 1}

	err := repo.Delete(ctx, model)
	assert.Error(t, err)
	assert.Equal(t, fakeError, err)
	repo.DB.Mock.AssertExpectations(t)
}
