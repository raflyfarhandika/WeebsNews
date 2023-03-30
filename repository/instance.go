package repository

import "sync"

type repositoryData struct {
	// UserRepository
	// NewsRepository
	CategoryRepository
	// CommentRepository
}

var repositoryInstance *repositoryData
var once sync.Once

func NewRepository() *repositoryData {
	return &repositoryData{
		// UserRepository:     NewUserRepository(),
		// NewsRepository:     NewNewsRepository(),
		CategoryRepository: NewCategoryRepository(),
		// CommentRepository:  NewCommentRepository(),
	}
}

func InitNewRepository() *repositoryData {
	once.Do(func() {
		repositoryInstance = NewRepository()
	})
	return repositoryInstance
}