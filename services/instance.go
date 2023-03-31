package services

import (
	"sync"
	"weebsnews/repository"
)

type serviceData struct {
	CategoryService
	UsersService
}

var repositoryInstance = repository.InitNewRepository()
var serviceInstance *serviceData
var once sync.Once

func NewService() *serviceData {
	return &serviceData{
		CategoryService:	NewCategoryService(repositoryInstance.CategoryRepository),
		UsersService:		NewUsersService(repositoryInstance.UsersRepository),
	}
}

func InitNewService() *serviceData {
	once.Do(func() {
		serviceInstance = NewService()
	})
	return serviceInstance
}