package services

import (
	"sync"
	"weebsnews/repository"
)

type serviceData struct {
	CategoryService
	UsersService
	NewsService
	CommentService
}

var repositoryInstance = repository.InitNewRepository()
var serviceInstance *serviceData
var once sync.Once

func NewService() *serviceData {
	return &serviceData{
		CategoryService:	NewCategoryService(repositoryInstance.CategoryRepository),
		UsersService:		NewUsersService(repositoryInstance.UsersRepository),
		NewsService:		NewNewsService(repositoryInstance.NewsRepository),
		CommentService:		NewCommentService(repositoryInstance.CommentRepository),
	}
}

func InitNewService() *serviceData {
	once.Do(func() {
		serviceInstance = NewService()
	})
	return serviceInstance
}