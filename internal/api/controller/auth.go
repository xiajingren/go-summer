package controller

import (
	"net/http"

	"github.com/xiajingren/go-summer/internal/api/dto"
	"github.com/xiajingren/go-summer/internal/api/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController() AuthController {
	return AuthController{authService: service.NewAuthService()}
}

func (ctl AuthController) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := ctl.authService.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ctl AuthController) Refresh(c *gin.Context) {
	var refreshRequest dto.RefreshRequest
	if err := c.ShouldBind(&refreshRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
