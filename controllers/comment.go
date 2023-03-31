package controllers

import (
	"net/http"
	"strconv"
	"weebsnews/model"
	"weebsnews/services"

	"github.com/gin-gonic/gin"
)

type CommentController interface {
	Create(c *gin.Context)
	GetCommentByNewsId(c *gin.Context)
	GetCommentById(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type commentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) CommentController {
	return &commentController{service}
}

func (controller *commentController) Create(c *gin.Context) {
	var request model.Comment
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := controller.service.Create(request)
	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H {
		"data" : response.Data,
		"message" : "Comment created successfully",
	})
}

func (controller *commentController) GetCommentByNewsId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request model.Comment
	request.NewsID = id

	response := controller.service.GetCommentByNewsId(request)
	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H {
		"data" : response.Data,
		"message" : "Comment fetched successfully",
	})
}

func (controller *commentController) GetCommentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := controller.service.GetCommentById(id)
	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H {
		"data" : response.Data,
		"message" : "Comment fetched successfully",
	})
}

func (controller *commentController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request model.Comment
	request.ID = id
	err = c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	search := controller.service.GetCommentById(request.ID)

	if search.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": search.Error,
		})
		return
	}

	response := controller.service.Update(request)
	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H {
		"message" : "Comment updated successfully",
	})
}

func (controller *commentController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request model.Comment
	request.ID = id

	search := controller.service.GetCommentById(int(request.ID))

	if search.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": search.Error,
		})
		return
	}

	response := controller.service.Delete(id)
	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H {
		"message" : "Comment deleted successfully",
	})
}