package service

import (
	"errors"

	"github.com/xiajingren/go-summer/internal/api/dto"
	"github.com/xiajingren/go-summer/pkg/utils"
	"github.com/xiajingren/go-summer/store"
)

type UserService struct {
	userRepository store.UserRepository
}

func NewUserService() UserService {
	return UserService{
		userRepository: store.NewUserRepository(),
	}
}

func (service UserService) Register(req dto.RegisterRequest) error {
	var err error

	exists, err := service.userRepository.Exists(req.Username)
	if err != nil {
		return errors.New("user register fail")
	}

	if exists {
		return errors.New("user already exists")
	}

	user := store.User{
		Username: req.Username,
		Password: utils.HashAndSalt([]byte(req.Password)),
	}
	err = service.userRepository.Create(&user)
	if err != nil {
		return errors.New("user register fail")
	}
	return nil
}
