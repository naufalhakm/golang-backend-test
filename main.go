package main

import (
	"golang-backend-test/database"
	"golang-backend-test/factory"
	"golang-backend-test/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewSQLiteConnection()
	if err != nil {
		panic(err)
	}
	router := gin.New()
	factory := factory.InitFactory(db)
	routes.NewRoutes(router, factory)
	router.Run(":8080")
}
