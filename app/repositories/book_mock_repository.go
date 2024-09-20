package repositories

import (
	"context"
	"golang-backend-test/app/models"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockBookRepository struct {
	mock.Mock
}

func (mock *MockBookRepository) FindBookById(ctx context.Context, db *gorm.DB, id int) (*models.Book, error) {
	args := mock.Called(ctx, db, id)
	if book, ok := args.Get(0).(*models.Book); ok {
		return book, args.Error(1)
	}
	return nil, args.Error(1)
}

func (mock *MockBookRepository) GetListBooks(ctx context.Context, db *gorm.DB) ([]*models.Book, error) {
	args := mock.Called(ctx, db)
	if books, ok := args.Get(0).([]*models.Book); ok {
		return books, args.Error(1)
	}
	return nil, args.Error(1)
}

func (mock *MockBookRepository) CreateBook(ctx context.Context, db *gorm.DB, book *models.Book) error {
	args := mock.Called(ctx, db, book)
	return args.Error(0)
}

func (mock *MockBookRepository) UpdateBook(ctx context.Context, db *gorm.DB, book *models.Book) error {
	args := mock.Called(ctx, db, book)
	return args.Error(0)
}

func (mock *MockBookRepository) DeleteBook(ctx context.Context, db *gorm.DB, id int) error {
	args := mock.Called(ctx, db, id)
	return args.Error(0)
}
