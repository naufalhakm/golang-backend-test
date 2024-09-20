package repositories

import (
	"context"
	"errors"
	"golang-backend-test/app/models"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAuthorById(ctx context.Context, db *gorm.DB, id int) (*models.Author, error)
	GetListAuthors(ctx context.Context, db *gorm.DB) ([]*models.Author, error)
	CreateAuthor(ctx context.Context, db *gorm.DB, author *models.Author) error
	UpdateAuthor(ctx context.Context, db *gorm.DB, author *models.Author) error
	DeleteAuthor(ctx context.Context, db *gorm.DB, id int) error
}

type AuthorRepositoryImpl struct {
}

func NewAuthorRepository() AuthorRepository {
	return &AuthorRepositoryImpl{}
}

func (repository *AuthorRepositoryImpl) FindAuthorById(ctx context.Context, db *gorm.DB, id int) (*models.Author, error) {
	var author models.Author
	if err := db.WithContext(ctx).First(&author, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("author not found")
		}
		return nil, err
	}
	return &author, nil
}
func (repository *AuthorRepositoryImpl) GetListAuthors(ctx context.Context, db *gorm.DB) ([]*models.Author, error) {
	var authors []*models.Author
	if err := db.WithContext(ctx).Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}
func (repository *AuthorRepositoryImpl) CreateAuthor(ctx context.Context, db *gorm.DB, author *models.Author) error {
	if err := db.WithContext(ctx).Create(author).Error; err != nil {
		return err
	}
	return nil
}
func (repository *AuthorRepositoryImpl) UpdateAuthor(ctx context.Context, db *gorm.DB, author *models.Author) error {
	if err := db.WithContext(ctx).Save(author).Error; err != nil {
		return err
	}
	return nil
}
func (repository *AuthorRepositoryImpl) DeleteAuthor(ctx context.Context, db *gorm.DB, id int) error {
	if err := db.WithContext(ctx).Delete(&models.Author{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("author not found")
		}
		return err
	}
	return nil
}
