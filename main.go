package main

import (
	"weebsnews/controllers"
	"weebsnews/database"
	"weebsnews/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	DB := database.DBStart()
	database.DBMigrate()
	defer DB.Close()

	router := gin.Default()
	controllerInstance := controllers.InitNewController()

	//Public API
	public := router.Group("/api/v1")
	public.POST("/login", controllerInstance.UsersController.Login)
	public.POST("/register", controllerInstance.UsersController.Register)
	public.GET("/news/:id/detail", controllerInstance.NewsController.GetAllNewsContentById)
	public.GET("/category", controllerInstance.CategoryController.GetAll)
	public.GET("/news", controllerInstance.NewsController.GetAllNews)
	public.GET("/comment/news/:id", controllerInstance.CommentController.GetCommentByNewsId)


	// Protected API
	protected := router.Group("/api/v1")
	protected.Use(middleware.JwtAuthMiddleware())
	
	protected.POST("/category", controllerInstance.CategoryController.Create)
	protected.GET("/category/:id", controllerInstance.CategoryController.GetByID)
	protected.PUT("/category/:id", controllerInstance.CategoryController.Update)
	protected.DELETE("/category/:id", controllerInstance.CategoryController.Delete)
	
	protected.POST("/users", controllerInstance.UsersController.Create)
	protected.GET("/users", controllerInstance.UsersController.GetAll)
	protected.GET("/users/:id", controllerInstance.UsersController.GetByID)
	protected.PUT("/users/:id", controllerInstance.UsersController.Update)
	protected.DELETE("/users/:id", controllerInstance.UsersController.Delete)

	protected.POST("/news", controllerInstance.NewsController.Create)
	protected.GET("/news/:id", controllerInstance.NewsController.GetNewsById)
	protected.PUT("/news/:id", controllerInstance.NewsController.Update)
	protected.DELETE("/news/:id", controllerInstance.NewsController.Delete)

	protected.POST("/comment", controllerInstance.CommentController.Create)
	protected.GET("/comment/:id", controllerInstance.CommentController.GetCommentById)
	protected.PUT("/comment/:id", controllerInstance.CommentController.Update)
	protected.DELETE("/comment/:id", controllerInstance.CommentController.Delete)
	
	
	router.Run("localhost:5000")

}
