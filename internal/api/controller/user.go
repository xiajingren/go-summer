package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiajingren/go-summer/internal/api/dto"
	"github.com/xiajingren/go-summer/internal/api/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController() UserController {
	return UserController{
		userService: service.NewUserService(),
	}
}

func (ctl UserController) Register(c *gin.Context) {
	var registerRequest dto.RegisterRequest
	if err := c.ShouldBind(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctl.userService.Register(registerRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (ctl UserController) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
