package middleware

import (
	"net/http"
	"strings"

	"github.com/xiajingren/go-summer/internal/api/consts"
	"github.com/xiajingren/go-summer/internal/api/model"
	"github.com/xiajingren/go-summer/internal/api/service"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authSchema := consts.AUTH_SCHEMA + " "

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, authSchema) {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		tokenString := authHeader[len(authSchema):]

		claims, valid, err := service.NewJWTService().ParseTokenWithClaims(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		currentUser := model.NewCurrentUser(claims.Username)
		c.Set(consts.CURRENT_USER, currentUser)

		c.Next()
	}
}
