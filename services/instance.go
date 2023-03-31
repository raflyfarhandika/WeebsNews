package services

import (
	"sync"
	"weebsnews/repository"
)

type serviceData struct {
	CategoryService
	UsersService
	NewsService
}

var repositoryInstance = repository.InitNewRepository()
var serviceInstance *serviceData
var once sync.Once

func NewService() *serviceData {
	return &serviceData{
		CategoryService:	NewCategoryService(repositoryInstance.CategoryRepository),
		UsersService:		NewUsersService(repositoryInstance.UsersRepository),
		NewsService:		NewNewsService(repositoryInstance.NewsRepository),
	}
}

func InitNewService() *serviceData {
	once.Do(func() {
		serviceInstance = NewService()
	})
	return serviceInstance
}