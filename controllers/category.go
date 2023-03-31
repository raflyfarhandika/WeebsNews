package controllers

import (
	"net/http"
	"strconv"
	"weebsnews/model"
	"weebsnews/services"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type categoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) CategoryController {
	return &categoryController{service}
}

func (controller *categoryController) Create(c *gin.Context) {
	var request model.Category
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	isValid, error := request.ValidateRequest()

	if !isValid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validation_error": error,
		})
		return
	}

	response := controller.service.Create(request)
	c.IndentedJSON(response.StatusCode, gin.H{
		"data" : response.Data,
		"message" : "Category created successfully",
	})
}

func (controller *categoryController) GetAll(c *gin.Context) {
	
	response := controller.service.GetAll()

	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H{
		"data" : response.Data,
		"message" : "Get All Category",
	})
}

func (controller *categoryController) GetByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := controller.service.GetByID(int64(userID))
	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H{
		"data" : response.Data,
		"message" : "Category list",
	})
}

func (controller *categoryController) Update(c *gin.Context) {
	var request model.Category
	c.ShouldBindJSON(&request)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	request.ID = id

	isValid, error := request.ValidateRequest()

	if !isValid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validation_error": error,
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

	c.IndentedJSON(response.StatusCode, gin.H{
		"message" : "Category updated successfully",
	})
}

func (controller *categoryController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request model.Category
	request.ID = id

	response := controller.service.Delete(request)
	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H{
		"message" : "Category deleted successfully",
	})
}