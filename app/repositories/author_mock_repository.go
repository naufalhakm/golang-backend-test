package repositories

import (
	"context"
	"golang-backend-test/app/models"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockAuthorRepository struct {
	mock.Mock
}

func (mock *MockAuthorRepository) FindAuthorById(ctx context.Context, db *gorm.DB, id int) (*models.Author, error) {
	args := mock.Called(ctx, db, id)
	if author, ok := args.Get(0).(*models.Author); ok {
		return author, args.Error(1)
	}
	return nil, args.Error(1)
}

func (mock *MockAuthorRepository) GetListAuthors(ctx context.Context, db *gorm.DB) ([]*models.Author, error) {
	args := mock.Called(ctx, db)
	if authors, ok := args.Get(0).([]*models.Author); ok {
		return authors, args.Error(1)
	}
	return nil, args.Error(1)
}

func (mock *MockAuthorRepository) CreateAuthor(ctx context.Context, db *gorm.DB, author *models.Author) error {
	args := mock.Called(ctx, db, author)
	return args.Error(0)
}

func (mock *MockAuthorRepository) UpdateAuthor(ctx context.Context, db *gorm.DB, author *models.Author) error {
	args := mock.Called(ctx, db, author)
	return args.Error(0)
}

func (mock *MockAuthorRepository) DeleteAuthor(ctx context.Context, db *gorm.DB, id int) error {
	args := mock.Called(ctx, db, id)
	return args.Error(0)
}
