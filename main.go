package main

import (
	"weebsnews/database"
)

func main() {

	DB := database.DBStart()
	database.DBMigrate()
	defer DB.Close()

	// router := gin.Default()
	// controllerInstance := controllers.InitNewController()

	// router.POST("/category", controllerInstance.CategoryController.Create)
	// router.GET("/category", controllerInstance.CategoryController.GetAll)
	// router.GET("/category/:id", controllerInstance.CategoryController.GetByID)
	// router.PUT("/category/:id", controllerInstance.CategoryController.Update)
	// router.DELETE("/category/:id", controllerInstance.CategoryController.Delete)
	
	
	
	// router.Run("localhost:5000")

}
