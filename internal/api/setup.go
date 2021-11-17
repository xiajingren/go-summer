package api

import (
	"net/http"

	"github.com/xiajingren/go-summer/internal/api/controller"
	"github.com/xiajingren/go-summer/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.New()

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, "MethodNotAllowed")
	})

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	registerRouter(r)

	r.Run()
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
	}

	auth_v1 := r.Group("/api/v1", middleware.JwtAuth())
	{
		user := auth_v1.Group("/users")
		{
			user.GET("", userController.GetUsers)
		}
	}
}
