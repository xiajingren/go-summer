package api

import (
	"net/http"

	"github.com/xiajingren/go-summer/conf"
	"github.com/xiajingren/go-summer/internal/api/controller"
	"github.com/xiajingren/go-summer/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(conf.Conf.Api.Gin_Mode)
	r := gin.New()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	})

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	registerRouter(r)

	r.Run(":" + conf.Conf.Api.Gin_Port)
}

var (
	authController controller.AuthController = controller.NewAuthController()
	userController controller.UserController = controller.NewUserController()
)

func registerRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authController.Login)
		}
		user := v1.Group("/users")
		{
			user.POST("/register", userController.Register)
		}
	}

	auth_v1 := r.Group("/api/v1", middleware.JwtAuth())
	{
		user := auth_v1.Group("/users")
		{
			user.GET("", userController.GetUsers)
		}
	}
}
