package response

import (
	"github.com/xiajingren/go-summer/conf"
	"github.com/xiajingren/go-summer/internal/api/consts"
	"github.com/xiajingren/go-summer/pkg/utils"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

func NewTokenResponse(access_token string) TokenResponse {
	return TokenResponse{
		AccessToken:  access_token,
		RefreshToken: utils.NewBase64RandomString(),
		TokenType:    consts.AUTH_SCHEMA,
		ExpiresIn:    conf.Conf.Api.Jwt_Exp,
	}
}
