package services

import (
	"weebsnews/model"
	"weebsnews/repository"
)

type NewsService interface {
	Create(request model.News) model.Response
	GetAllNews() model.Response
	GetNewsById(id int) model.Response
	Update(request model.News) model.Response
	Delete(id int) model.Response
}

type newsService struct {
	repo repository.NewsRepository
}

func NewNewsService(repo repository.NewsRepository) NewsService {
	return &newsService{repo}
}

func (service *newsService) Create(news model.News) model.Response {
	result, err := service.repo.Create(news)

	if err != nil {
		return model.Response{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	return model.Response{
		Error:      "",
		StatusCode: 201,
		Data: result,
	}
}

func (service *newsService) GetAllNews() model.Response {
	result, err := service.repo.GetAllNews()

	if err != nil {
		return model.Response{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	return model.Response{
		Error:      "",
		StatusCode: 200,
		Data:       result,
	}
}

func (service *newsService) GetNewsById(id int) model.Response {
	result, err := service.repo.GetNewsById(id)

	if err != nil {
		return model.Response{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	return model.Response{
		Error:      "",
		StatusCode: 200,
		Data:       result,
	}
}

func (service *newsService) Update(news model.News) model.Response {
	err := service.repo.Update(news)

	if err != nil {
		return model.Response{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	return model.Response{
		Error:      "",
		StatusCode: 200,
		Data:       map[string]interface{}{},
	}
}

func (service *newsService) Delete(id int) model.Response {
	err := service.repo.Delete(id)

	if err != nil {
		return model.Response{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	return model.Response{
		Error:      "",
		StatusCode: 200,
		Data:       map[string]interface{}{},
	}
}