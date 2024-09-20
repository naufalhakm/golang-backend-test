package repositories

import (
	"context"
	"golang-backend-test/app/models"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	mock.Mock
}

func (mock *MockUserRepository) CreateUser(ctx context.Context, db *gorm.DB, user *models.User) error {
	args := mock.Called(ctx, db, user)
	return args.Error(0)
}

func (mock *MockUserRepository) FindUserByUsername(ctx context.Context, db *gorm.DB, username string) (*models.User, error) {
	args := mock.Called(ctx, db, username)
	if user, ok := args.Get(0).(*models.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (mock *MockUserRepository) HashPassword(password string) (string, error) {
	args := mock.Called(password)
	return args.String(0), args.Error(1)
}
