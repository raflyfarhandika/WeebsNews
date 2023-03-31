package controllers

import (
	"sync"
	"weebsnews/services"
)

type controllerData struct {
	CategoryController
	UsersController
	NewsController
}

var serviceInstance = services.InitNewService()
var controllerInstance *controllerData
var once sync.Once

func NewController() *controllerData {
	return &controllerData{
		CategoryController: NewCategoryController(serviceInstance.CategoryService),
		UsersController:    NewUsersController(serviceInstance.UsersService),
		NewsController:     NewNewsController(serviceInstance.NewsService),
	}
}

func InitNewController() *controllerData {
	once.Do(func() {
		controllerInstance = NewController()
	})
	return controllerInstance
}