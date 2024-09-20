package factory

import (
	"golang-backend-test/app/controllers"
	"golang-backend-test/app/repositories"
	"golang-backend-test/app/services"

	"gorm.io/gorm"
)

type Provider struct {
	UserProvider   controllers.UserController
	BookProvider   controllers.BookController
	AuthorProvider controllers.AuthorController
}

func InitFactory(db *gorm.DB) *Provider {

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo, db)
	userController := controllers.NewUserController(userService)

	bookRepo := repositories.NewBookRepository()
	authorRepo := repositories.NewAuthorRepository()
	bookService := services.NewBookService(bookRepo, authorRepo, db)
	bookController := controllers.NewBookController(bookService)

	authorService := services.NewAuthorService(authorRepo, db)
	authorController := controllers.NewAuthorController(authorService)

	return &Provider{
		UserProvider:   userController,
		BookProvider:   bookController,
		AuthorProvider: authorController,
	}
}
