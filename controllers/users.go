package controllers

import (
	"net/http"
	"strconv"
	"weebsnews/model"
	"weebsnews/services"

	"github.com/gin-gonic/gin"
)

type UsersController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type usersController struct {
	service services.UsersService
}

func NewUsersController(service services.UsersService) UsersController {
	return &usersController{service}
}

func (controller *usersController) Create(c *gin.Context) {
	var request model.Users
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
		"message" : "User created successfully",
	})
}

func (controller *usersController) GetAll(c *gin.Context) {
	
	response := controller.service.GetAll()

	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H{
		"data" : response.Data,
		"message" : "All users",
	})
}

func (controller *usersController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := controller.service.GetByID(int(id))
	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H{
		"data" : response.Data,
		"message" : "User by ID",
	})
}

func (controller *usersController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request model.Users
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

	search := controller.service.GetByID(int(request.ID))

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
		"message" : "User updated successfully",
	})
}

func (controller *usersController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request model.Users
	request.ID = id

	search := controller.service.GetByID(int(request.ID))

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
		"message" : "User deleted successfully",
	})
}

func (controller *usersController) Login(c *gin.Context) {
	var request model.UsersLogin
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := controller.service.Login(request)
	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H{
		"data" : response.Data,
		"message" : "User logged in successfully",
	})
}

func (controller *usersController) Register(c *gin.Context) {
	var request model.Users
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

	response := controller.service.Register(request)
	if response.Error != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": response.Error,
		})
		return
	}

	c.IndentedJSON(response.StatusCode, gin.H{
		"data" : response.Data,
		"message" : "User registered successfully",
	})
}