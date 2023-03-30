package services

import (
	"weebsnews/model"
	"weebsnews/repository"
)

type CategoryService interface {
	Create(request model.Category) model.Response
	GetAll() model.Response
	GetByID(request model.Category) model.Response
	Update(request model.Category) model.Response
	Delete(request model.Category) model.Response
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (service *categoryService) Create(request model.Category) model.Response {
	result, err := service.repo.Create(request)

	if err != nil {
		return model.Response{
			Error: 		err.Error(),
			StatusCode: 500,
			Data: 		map[string]interface{}{},
		}
	}

	return model.Response{
		Error: 		"",
		StatusCode: 201,
		Data: 		map[string]interface{}{
			"category": result,
		},
	}

}

func (service *categoryService) GetAll() model.Response {
	result, err := service.repo.GetAll()

	if err != nil {
		return model.Response{
			Error: 		err.Error(),
			StatusCode: 500,
			Data: 		map[string]interface{}{},
		}
	}

	return model.Response{
		Error: 		"",
		StatusCode: 200,
		Data: 		map[string]interface{}{
			"categories": result,
		},
	}
}

func (service *categoryService) GetByID(request model.Category) model.Response {
	result, err := service.repo.GetByID(request)

	if err != nil {
		return model.Response{
			Error: 		err.Error(),
			StatusCode: 500,
			Data: 		map[string]interface{}{},
		}
	}

	return model.Response{
		Error: 		"",
		StatusCode: 200,
		Data: 		map[string]interface{}{
			"category": result,
		},
	}
}

func (service *categoryService) Update(request model.Category) model.Response {
	result, err := service.repo.GetByID(request)

	if err != nil {
		return model.Response{
			Error: 		err.Error(),
			StatusCode: 500,
			Data: 		map[string]interface{}{},
		}
	}

	if result.ID == 0 {
		return model.Response{
			Error: 		"Category not found",
			StatusCode: 404,
			Data: 		map[string]interface{}{},
		}
	}

	err = service.repo.Update(request)

	if err != nil {
		return model.Response{
			Error: 		err.Error(),
			StatusCode: 500,
			Data: 		map[string]interface{}{},
		}
	}

	return model.Response{
		Error: 		"",
		StatusCode: 200,
		Data: 		map[string]interface{}{
			"category": result,
		},
	}
}

func (service *categoryService) Delete(request model.Category) model.Response {
	err := service.repo.Delete(request)

	if err != nil {
		return model.Response{
			Error: 		err.Error(),
			StatusCode: 500,
			Data: 		map[string]interface{}{},
		}
	}

	return model.Response{
		Error: 		"",
		StatusCode: 200,
	}
}