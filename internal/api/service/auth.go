package service

import (
	"errors"

	"github.com/xiajingren/go-summer/internal/api/dto"
	"github.com/xiajingren/go-summer/store"
)

type AuthService struct {
	userRepository store.UserRepository
	jwtService     JWTService
}

func NewAuthService() AuthService {
	return AuthService{
		userRepository: store.NewUserRepository(),
		jwtService:     NewJWTService(),
	}
}

func (service AuthService) Login(req dto.LoginRequest) (*dto.TokenResponse, error) {
	var err error

	user, err := service.userRepository.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("wrong user name or password")
	}

	if user == nil {
		return nil, errors.New("wrong user name or password")
	}

	if user.Password != req.Password {
		return nil, errors.New("wrong user name or password")
	}

	token, err := service.jwtService.GenerateToken(req.Username)
	if err != nil {
		return nil, errors.New("an error occurred while generating the token")
	}

	return dto.NewTokenResponse(token), nil
}
