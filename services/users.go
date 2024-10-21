package services

import (
	"context"
	"github.com/SamPariatIL/weather-wrapper/entities"
	"github.com/SamPariatIL/weather-wrapper/repository"
	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(user *entities.UserDetails) (*string, error)
	UpdateUser(uid string, user *entities.UserDetails) (*string, error)
	DeleteUser(uid string) (*string, error)
	GenerateToken(uid string) (*string, error)
	SendVerificationEmail(email string) (*string, error)
	ResetPassword(email string) (*string, error)
}

type userService struct {
	userRepo repository.UserRepository
	logger   *zap.Logger
}

func NewUserService(ur repository.UserRepository, zl *zap.Logger) UserService {
	return &userService{
		userRepo: ur,
		logger:   zl,
	}
}

func (us *userService) CreateUser(user *entities.UserDetails) (*string, error) {
	userId, err := us.userRepo.CreateUser(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return userId, nil
}

func (us *userService) UpdateUser(uid string, user *entities.UserDetails) (*string, error) {
	userId, err := us.userRepo.UpdateUser(context.Background(), uid, user)
	if err != nil {
		return nil, err
	}

	return userId, nil
}

func (us *userService) DeleteUser(uid string) (*string, error) {
	userId, err := us.userRepo.DeleteUser(context.Background(), uid)
	if err != nil {
		return nil, err
	}

	return userId, nil
}

func (us *userService) GenerateToken(uid string) (*string, error) {
	token, err := us.userRepo.GenerateToken(context.Background(), uid)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (us *userService) SendVerificationEmail(email string) (*string, error) {
	link, err := us.userRepo.SendVerificationEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (us *userService) ResetPassword(email string) (*string, error) {
	link, err := us.userRepo.ResetPassword(context.Background(), email)
	if err != nil {
		return nil, err
	}

	return link, nil
}
