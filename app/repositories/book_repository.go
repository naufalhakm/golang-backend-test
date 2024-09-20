package repositories

import (
	"context"
	"errors"
	"golang-backend-test/app/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindBookById(ctx context.Context, db *gorm.DB, id int) (*models.Book, error)
	GetListBooks(ctx context.Context, db *gorm.DB) ([]*models.Book, error)
	CreateBook(ctx context.Context, db *gorm.DB, book *models.Book) error
	UpdateBook(ctx context.Context, db *gorm.DB, book *models.Book) error
	DeleteBook(ctx context.Context, db *gorm.DB, id int) error
}

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (repositories *BookRepositoryImpl) FindBookById(ctx context.Context, db *gorm.DB, id int) (*models.Book, error) {
	var book models.Book
	if err := db.WithContext(ctx).Preload("Author").First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return &book, nil
}
func (repositories *BookRepositoryImpl) GetListBooks(ctx context.Context, db *gorm.DB) ([]*models.Book, error) {
	var books []*models.Book
	if err := db.WithContext(ctx).Preload("Author").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
func (repositories *BookRepositoryImpl) CreateBook(ctx context.Context, db *gorm.DB, book *models.Book) error {
	if err := db.WithContext(ctx).Create(book).Error; err != nil {
		return err
	}
	return nil
}
func (repositories *BookRepositoryImpl) UpdateBook(ctx context.Context, db *gorm.DB, book *models.Book) error {
	if err := db.WithContext(ctx).Save(book).Error; err != nil {
		return err
	}
	return nil
}
func (repositories *BookRepositoryImpl) DeleteBook(ctx context.Context, db *gorm.DB, id int) error {
	if err := db.WithContext(ctx).Delete(&models.Book{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("book not found")
		}
		return err
	}
	return nil
}
