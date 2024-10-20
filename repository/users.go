package repository

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/SamPariatIL/weather-wrapper/entities"
	"go.uber.org/zap"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.UserDetails) (*string, error)
	UpdateUser(ctx context.Context, uid string, user *entities.UserDetails) (*string, error)
	DeleteUser(ctx context.Context, uid string) (*string, error)
	GenerateToken(ctx context.Context, uid string) (*string, error)
	SendVerificationEmail(ctx context.Context, email string) (*string, error)
	ResetPassword(ctx context.Context, email string) (*string, error)
}

type userRepository struct {
	firebaseAuth *auth.Client
	logger       *zap.Logger
}

func NewUserRepository(fa *auth.Client, zl *zap.Logger) UserRepository {
	return &userRepository{
		firebaseAuth: fa,
		logger:       zl,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *entities.UserDetails) (*string, error) {
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		EmailVerified(user.EmailVerified).
		PhoneNumber(user.PhoneNumber).
		Password(user.Password).
		DisplayName(user.DisplayName).
		Disabled(user.Disabled)

	if user.PhotoURL != nil {
		params.PhotoURL(*user.PhotoURL)
	}

	createdUser, err := ur.firebaseAuth.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	ur.logger.Info(fmt.Sprintf("created user %s", createdUser.UID))
	return &createdUser.UID, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, uid string, user *entities.UserDetails) (*string, error) {
	params := (&auth.UserToUpdate{}).
		Email(user.Email).
		EmailVerified(user.EmailVerified).
		PhoneNumber(user.PhoneNumber).
		Password(user.Password).
		DisplayName(user.DisplayName).
		Disabled(user.Disabled)

	if user.PhotoURL != nil {
		params.PhotoURL(*user.PhotoURL)
	}

	updatedUser, err := ur.firebaseAuth.UpdateUser(ctx, uid, params)
	if err != nil {
		return nil, err
	}

	ur.logger.Info(fmt.Sprintf("updated user %s", updatedUser.UID))
	return &uid, nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, uid string) (*string, error) {
	err := ur.firebaseAuth.DeleteUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	ur.logger.Info(fmt.Sprintf("deleted user %s", uid))
	return &uid, nil
}

func (ur *userRepository) GenerateToken(ctx context.Context, uid string) (*string, error) {
	token, err := ur.firebaseAuth.CustomToken(ctx, uid)
	if err != nil {
		return nil, err
	}

	ur.logger.Info(fmt.Sprintf("logging in user %s - token %s", uid, token))
	return &uid, nil
}

func (ur *userRepository) SendVerificationEmail(ctx context.Context, email string) (*string, error) {
	link, err := ur.firebaseAuth.EmailVerificationLink(ctx, email)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (ur *userRepository) ResetPassword(ctx context.Context, email string) (*string, error) {
	link, err := ur.firebaseAuth.PasswordResetLink(ctx, email)
	if err != nil {
		return nil, err
	}

	return &link, nil
}
