package main

import (
	"weebsnews/controllers"
	"weebsnews/database"

	"github.com/gin-gonic/gin"
)

func main() {

	DB := database.DBStart()
	database.DBMigrate()
	defer DB.Close()

	router := gin.Default()
	controllerInstance := controllers.InitNewController()

	router.POST("/category", controllerInstance.CategoryController.Create)
	router.GET("/category", controllerInstance.CategoryController.GetAll)
	router.GET("/category/:id", controllerInstance.CategoryController.GetByID)
	router.PUT("/category/:id", controllerInstance.CategoryController.Update)
	router.DELETE("/category/:id", controllerInstance.CategoryController.Delete)
	
	router.POST("/users", controllerInstance.UsersController.Create)
	router.GET("/users", controllerInstance.UsersController.GetAll)
	router.GET("/users/:id", controllerInstance.UsersController.GetByID)
	router.PUT("/users/:id", controllerInstance.UsersController.Update)
	router.DELETE("/users/:id", controllerInstance.UsersController.Delete)
	
	router.Run("localhost:5000")

}
