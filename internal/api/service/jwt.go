package service

import (
	"errors"
	"time"

	"github.com/xiajingren/go-summer/conf"
	"github.com/xiajingren/go-summer/pkg/utils"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService struct{}

func NewJWTService() JWTService {
	return JWTService{}
}

type JWTClaims struct {
	Username string
	jwt.StandardClaims
}

func newJWTClaims(username string) (claim JWTClaims) {
	claim.Username = username
	claim.IssuedAt = time.Now().Unix()
	claim.ExpiresAt = time.Now().Unix() + conf.Conf.Api.Jwt_Exp
	claim.NotBefore = time.Now().Unix()
	claim.Id = utils.NewUUIdString()
	return
}

func (service JWTService) GenerateToken(username string) (string, error) {
	claim := newJWTClaims(username)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(conf.Conf.Api.Jwt_Key)
}

func (service JWTService) ParseTokenWithClaims(tokenString string) (*JWTClaims, bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return conf.Conf.Api.Jwt_Key, nil
	})

	if err != nil {
		return nil, false, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok {
		return claims, token.Valid, nil
	}

	return nil, false, errors.New("an error occurred while parse with the claims")
}
