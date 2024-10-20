package handlers

import (
	"context"
	"github.com/SamPariatIL/weather-wrapper/entities"
	"github.com/SamPariatIL/weather-wrapper/services"
	"github.com/SamPariatIL/weather-wrapper/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserHandler interface {
	CreateUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
	GenerateToken(ctx *fiber.Ctx) error
	SendVerificationEmail(ctx *fiber.Ctx) error
	ResetPassword(ctx *fiber.Ctx) error
}

type userHandler struct {
	userService services.UserService
	logger      *zap.Logger
}

func NewUserHandler(us services.UserService, zl *zap.Logger) UserHandler {
	return &userHandler{
		userService: us,
		logger:      zl,
	}
}

// CreateUser godoc
// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body entities.UserDetails true "User details"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /users/signup [post]
func (uh *userHandler) CreateUser(ctx *fiber.Ctx) error {
	user := new(entities.UserDetails)

	if err := ctx.BodyParser(user); err != nil {
		uh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.CustomResponse(nil, fiber.StatusBadRequest, "", err.Error()))
	}

	uid, err := uh.userService.CreateUser(context.Background(), user)
	if err != nil {
		uh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(utils.CustomResponse(nil, fiber.StatusInternalServerError, userCreationError, err.Error()))
	}

	user.UID = uid

	uh.logger.Info(successCreatingUser)
	return ctx.Status(fiber.StatusCreated).
		JSON(utils.CustomResponse(user, fiber.StatusCreated, "", successCreatingUser))
}

// UpdateUser godoc
// @Summary Update user
// @Description Update a user
// @Tags users
// @Accept json
// @Produce json
// @Param uid path string true "User ID"
// @Param user body entities.UserDetails true "User details"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /users/{uid} [put]
func (uh *userHandler) UpdateUser(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")
	user := new(entities.UserDetails)

	if err := ctx.BodyParser(user); err != nil {
		uh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.CustomResponse(nil, fiber.StatusBadRequest, "", err.Error()))
	}

	updatedUserId, err := uh.userService.UpdateUser(context.Background(), uid, user)
	if err != nil {
		uh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(utils.CustomResponse(nil, fiber.StatusInternalServerError, userUpdationError, err.Error()))
	}

	uh.logger.Info(successUpdatingUser)
	return ctx.Status(fiber.StatusOK).
		JSON(utils.CustomResponse(updatedUserId, fiber.StatusOK, "", successUpdatingUser))
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user
// @Tags users
// @Accept json
// @Produce json
// @Param uid path string true "User ID"
// @Success 200
// @Failure 500
// @Router /users/{uid} [delete]
func (uh *userHandler) DeleteUser(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")

	deletedUserId, err := uh.userService.DeleteUser(context.Background(), uid)
	if err != nil {
		uh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(utils.CustomResponse(nil, fiber.StatusInternalServerError, userDeletionError, err.Error()))
	}

	uh.logger.Info(successDeletingUser)
	return ctx.Status(fiber.StatusOK).
		JSON(utils.CustomResponse(deletedUserId, fiber.StatusOK, "", successDeletingUser))
}

// GenerateToken godoc
// @Summary Generate token
// @Description Generate a token
// @Tags users
// @Accept json
// @Produce json
// @Param token body entities.UidBody true "Token body"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /users/token [post]
func (uh *userHandler) GenerateToken(ctx *fiber.Ctx) error {
	user := new(entities.UidBody)

	if err := ctx.BodyParser(user); err != nil {
		uh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.CustomResponse(nil, fiber.StatusBadRequest, "", err.Error()))
	}

	token, err := uh.userService.GenerateToken(context.Background(), user.UID)
	if err != nil {
		uh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(utils.CustomResponse(nil, fiber.StatusInternalServerError, tokenGenerationError, err.Error()))
	}

	uh.logger.Info(successGeneratingToken)
	return ctx.Status(fiber.StatusOK).
		JSON(utils.CustomResponse(token, fiber.StatusOK, "", successGeneratingToken))
}

// SendVerificationEmail godoc
// @Summary Send verification email
// @Description Send a verification email
// @Tags users
// @Accept json
// @Produce json
// @Param email body entities.EmailBody true "Email body"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /users/verify [post]
func (uh *userHandler) SendVerificationEmail(ctx *fiber.Ctx) error {
	emailBody := new(entities.EmailBody)

	if err := ctx.BodyParser(emailBody); err != nil {
		uh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.CustomResponse(nil, fiber.StatusBadRequest, "", err.Error()))
	}

	link, err := uh.userService.SendVerificationEmail(context.Background(), emailBody.Email)
	if err != nil {
		uh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(utils.CustomResponse(nil, fiber.StatusInternalServerError, emailSendingError, err.Error()))
	}

	uh.logger.Info(successSendingEmail)
	return ctx.Status(fiber.StatusOK).
		JSON(utils.CustomResponse(map[string]string{"link": *link}, fiber.StatusOK, "", successSendingEmail))
}

// ResetPassword godoc
// @Summary Reset password
// @Description Send a verification email to reset password
// @Tags users
// @Accept json
// @Produce json
// @Param email body entities.EmailBody true "Email body"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /users/reset-password [post]
func (uh *userHandler) ResetPassword(ctx *fiber.Ctx) error {
	emailBody := new(entities.EmailBody)

	if err := ctx.BodyParser(emailBody); err != nil {
		uh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.CustomResponse(nil, fiber.StatusBadRequest, "", err.Error()))
	}

	link, err := uh.userService.ResetPassword(context.Background(), emailBody.Email)
	if err != nil {
		uh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(utils.CustomResponse(nil, fiber.StatusInternalServerError, emailSendingError, err.Error()))
	}

	uh.logger.Info(successSendingEmail)
	return ctx.Status(fiber.StatusOK).
		JSON(utils.CustomResponse(map[string]string{"link": *link}, fiber.StatusOK, "", successSendingEmail))
}
