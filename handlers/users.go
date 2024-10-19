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
