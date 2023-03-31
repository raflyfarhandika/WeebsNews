package services

import (
	"weebsnews/model"
	"weebsnews/repository"
)

type CommentService interface {
	Create(request model.Comment) model.Response
	GetCommentByNewsId(request model.Comment) model.Response
	GetCommentById(id int) model.Response
	Update(request model.Comment) model.Response
	Delete(id int) model.Response
}

type commentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentService{repo}
}

func (service *commentService) Create(comment model.Comment) model.Response {
	result, err := service.repo.Create(comment)

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

func (service *commentService) GetCommentByNewsId(request model.Comment) model.Response {
	result, err := service.repo.GetCommentByNewsId(request)

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

func (service *commentService) GetCommentById(id int) model.Response {
	result, err := service.repo.GetCommentById(id)

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

func (service *commentService) Update(comment model.Comment) model.Response {
	err := service.repo.Update(comment)

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

func (service *commentService) Delete(id int) model.Response {
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