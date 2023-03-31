package services

import (
	"weebsnews/model"
	"weebsnews/repository"
)

type UsersService interface {
	Create(request model.Users) model.Response
	GetAll() model.Response
	GetByID(id int64) model.Response
	Update(request model.Users) model.Response
	Delete(request model.Users) model.Response
}

type usersService struct {
	repo repository.UsersRepository
}

func NewUsersService(repo repository.UsersRepository) UsersService {
	return &usersService{repo}
}

func (service *usersService) Create(users model.Users) model.Response {
	result, err := service.repo.Create(users)

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

func (service *usersService) GetAll() model.Response {
	result, err := service.repo.GetAll()

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

func (service *usersService) GetByID(id int64) model.Response {
	result, err := service.repo.GetByID(id)

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

func (service *usersService) Update(users model.Users) model.Response {
	err := service.repo.Update(users)

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

func (service *usersService) Delete(users model.Users) model.Response {
	err := service.repo.Delete(users)

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

