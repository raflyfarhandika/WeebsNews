package controllers

import (
	"net/http"
	"strconv"
	"weebsnews/model"
	"weebsnews/services"

	"github.com/gin-gonic/gin"
)

type NewsController interface {
	Create(c *gin.Context)
	GetAllNews(c *gin.Context)
	GetNewsById(c *gin.Context)
	GetAllNewsContentById (c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type newsController struct {
	service services.NewsService
}

func NewNewsController(service services.NewsService) NewsController{
	return &newsController{service}
}

func (controller *newsController) Create(c *gin.Context) {
	var request model.News
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
	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H{
		"data" : response.Data,
		"message" : "News created successfully",
	})	
}

func (controller *newsController) GetAllNews(c *gin.Context) {
	response := controller.service.GetAllNews()

	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H{
		"data" : response.Data,
		"message" : "News fetched successfully",
	})
}

func (controller *newsController) GetNewsById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := controller.service.GetNewsById(int(id))
	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H{
		"data" : response.Data,
		"message" : "News fetched successfully",
	})
}

func (controller *newsController) GetAllNewsContentById (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := controller.service.GetNewsWithCategoriesAndCommentsByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"data" : response,
		"message" : "News fetched successfully",
	})
}

func (controller *newsController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request model.News
	request.ID = id
	err = c.ShouldBindJSON(&request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// isValid, error := request.ValidateRequest()

	// if !isValid {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"validation_error": error,
	// 	})
	// 	return
	// }

	search := controller.service.GetNewsById(int(request.ID))

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

	c.IndentedJSON(response.StatusCode, gin.H{
		"message" : "News updated successfully",
	})
}

func (controller *newsController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request model.News
	request.ID = id

	search := controller.service.GetNewsById(int(request.ID))

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

	c.IndentedJSON(response.StatusCode, gin.H{
		"message" : "News deleted successfully",
	})
}