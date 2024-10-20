package services

import (
	"context"
	"github.com/SamPariatIL/weather-wrapper/entities"
	"github.com/SamPariatIL/weather-wrapper/repository"
	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(ctx context.Context, user *entities.UserDetails) (*string, error)
	UpdateUser(ctx context.Context, uid string, user *entities.UserDetails) (*string, error)
	DeleteUser(ctx context.Context, uid string) (*string, error)
	GenerateToken(ctx context.Context, uid string) (*string, error)
	SendVerificationEmail(ctx context.Context, email string) (*string, error)
	ResetPassword(ctx context.Context, email string) (*string, error)
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

func (us *userService) CreateUser(ctx context.Context, user *entities.UserDetails) (*string, error) {
	userId, err := us.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return userId, nil
}

func (us *userService) UpdateUser(ctx context.Context, uid string, user *entities.UserDetails) (*string, error) {
	userId, err := us.userRepo.UpdateUser(ctx, uid, user)
	if err != nil {
		return nil, err
	}

	return userId, nil
}

func (us *userService) DeleteUser(ctx context.Context, uid string) (*string, error) {
	userId, err := us.userRepo.DeleteUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	return userId, nil
}

func (us *userService) GenerateToken(ctx context.Context, uid string) (*string, error) {
	token, err := us.userRepo.GenerateToken(ctx, uid)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (us *userService) SendVerificationEmail(ctx context.Context, email string) (*string, error) {
	link, err := us.userRepo.SendVerificationEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (us *userService) ResetPassword(ctx context.Context, email string) (*string, error) {
	link, err := us.userRepo.ResetPassword(ctx, email)
	if err != nil {
		return nil, err
	}

	return link, nil
}
