package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUsers(*gin.Context)
}

func NewUserController() UserController {
	return &userController{}
}

type userController struct{}

func (controller userController) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user": "admin",
	})
}
