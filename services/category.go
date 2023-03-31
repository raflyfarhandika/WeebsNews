package services

import (
	"weebsnews/model"
	"weebsnews/repository"
)

type CategoryService interface {
	Create(request model.Category) model.Response
	GetAll() model.Response
	GetByID(id int64) model.Response
	Update(request model.Category) model.Response
	Delete(request model.Category) model.Response
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (service *categoryService) Create(category model.Category) model.Response {
	result, err := service.repo.Create(category)

	if err != nil{
		return model.Response{
			Error: err.Error(),
			StatusCode: 500,
			Data: map[string]interface{}{},
		}
	}

	return model.Response{
		Error: "",
		StatusCode: 201,
		Data: map[string]interface{}{
			"category": result,
		},
	}
}

func (service *categoryService) GetAll() model.Response {
	result, err := service.repo.GetAll()

	if err != nil{
		return model.Response{
			Error: err.Error(),
			StatusCode: 500,
			Data: map[string]interface{}{},
		}
	}

	return model.Response{
		Error: "",
		StatusCode: 200,
		Data: result,
	}
}

func (service *categoryService) GetByID(id int64) model.Response {
	result, err := service.repo.GetByID(id)

	if err != nil{
		return model.Response{
			Error: err.Error(),
			StatusCode: 500,
			Data: map[string]interface{}{},
		}
	}

	return model.Response{
		Error: "",
		StatusCode: 200,
		Data: result,
	}
}

func (service *categoryService) Update(category model.Category) model.Response {
	err := service.repo.Update(category)

	if err != nil{
		return model.Response{
			Error: err.Error(),
			StatusCode: 500,
			Data: map[string]interface{}{},
		}
	}

	return model.Response{
		Error: "",
		StatusCode: 200,
		Data: map[string]interface{}{},
	}
}

func (service *categoryService) Delete(category model.Category) model.Response {
	err := service.repo.Delete(category)

	if err != nil{
		return model.Response{
			Error: err.Error(),
			StatusCode: 500,
			Data: map[string]interface{}{},
		}
	}

	return model.Response{
		Error: "",
		StatusCode: 200,
		Data: map[string]interface{}{},
	}
}