package repositories

import (
	"context"
	"errors"
	"golang-backend-test/app/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByUsername(ctx context.Context, db *gorm.DB, username string) (*models.User, error)
	CreateUser(ctx context.Context, db *gorm.DB, user *models.User) error
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repositories *UserRepositoryImpl) FindUserByUsername(ctx context.Context, db *gorm.DB, username string) (*models.User, error) {
	var user models.User
	if err := db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
func (repositories *UserRepositoryImpl) CreateUser(ctx context.Context, db *gorm.DB, user *models.User) error {
	if err := db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}
