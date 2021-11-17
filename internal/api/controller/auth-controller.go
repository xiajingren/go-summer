package controller

import (
	"net/http"

	"github.com/xiajingren/go-summer/internal/api/model/request"
	"github.com/xiajingren/go-summer/internal/api/model/response"
	"github.com/xiajingren/go-summer/internal/api/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(*gin.Context)
	RefreshToken(*gin.Context)
}

func NewAuthController() AuthController {
	return &authController{jwtService: service.NewJWTService()}
}

type authController struct {
	jwtService service.JWTService
}

func (controller authController) Login(c *gin.Context) {
	var LoginRequest request.LoginRequest
	if err := c.ShouldBind(&LoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// todo:
	if !(LoginRequest.Username == "admin" && LoginRequest.Password == "123456") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password error"})
		return
	}

	tokenString, err := controller.jwtService.GenerateToken(LoginRequest.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.NewTokenResponse(tokenString))
}

func (controller authController) RefreshToken(c *gin.Context) {
	var refreshRequest request.RefreshRequest
	if err := c.ShouldBind(&refreshRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
