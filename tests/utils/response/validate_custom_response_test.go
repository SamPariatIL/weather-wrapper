package tests

import (
	"github.com/SamPariatIL/weather-wrapper/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ValidateCustomResponseSuite struct {
	suite.Suite
}

func (suite *ValidateCustomResponseSuite) TestValidCity() {
	ok := "ok"
	created := "created"
	badRequest := "bad request"
	unauthorized := "unauthorized"
	forbidden := "forbidden"
	notFound := "not found"
	conflict := "conflict"
	internalServerError := "internal server error"
	somethingWentWrong := "something went wrong"
	noError := ""

	customResponses := []struct {
		data    any
		status  int
		error   string
		message string
	}{
		{data: "ok", status: 200, error: noError, message: ok},
		{data: "error", status: 201, error: noError, message: created},
		{data: "error", status: 400, error: somethingWentWrong, message: badRequest},
		{data: "error", status: 401, error: somethingWentWrong, message: unauthorized},
		{data: "error", status: 403, error: somethingWentWrong, message: forbidden},
		{data: "error", status: 404, error: somethingWentWrong, message: notFound},
		{data: "error", status: 409, error: somethingWentWrong, message: conflict},
		{data: "error", status: 500, error: somethingWentWrong, message: internalServerError},
		{data: fiber.Map{"message": ok}, status: 200, error: noError, message: ok},
		{data: fiber.Map{"message": created}, status: 201, error: noError, message: created},
		{data: fiber.Map{"message": badRequest}, status: 400, error: somethingWentWrong, message: badRequest},
		{data: fiber.Map{"message": unauthorized}, status: 401, error: somethingWentWrong, message: unauthorized},
		{data: fiber.Map{"message": forbidden}, status: 403, error: somethingWentWrong, message: forbidden},
		{data: fiber.Map{"message": notFound}, status: 404, error: somethingWentWrong, message: notFound},
		{data: fiber.Map{"message": conflict}, status: 409, error: somethingWentWrong, message: conflict},
		{data: fiber.Map{"message": internalServerError}, status: 500, error: somethingWentWrong, message: internalServerError},
		{data: nil, status: 200, error: noError, message: ok},
		{data: nil, status: 201, error: noError, message: created},
		{data: nil, status: 400, error: somethingWentWrong, message: badRequest},
		{data: nil, status: 401, error: somethingWentWrong, message: unauthorized},
		{data: nil, status: 403, error: somethingWentWrong, message: forbidden},
		{data: nil, status: 404, error: somethingWentWrong, message: notFound},
		{data: nil, status: 409, error: somethingWentWrong, message: conflict},
		{data: nil, status: 500, error: somethingWentWrong, message: internalServerError},
	}

	for _, customResponse := range customResponses {
		response := utils.CustomResponse(customResponse.data, customResponse.status, customResponse.error, customResponse.message)
		suite.Equal(customResponse.data, response["data"])
		suite.Equal(customResponse.status, response["status"])
		suite.Equal(customResponse.error, response["error"])
		suite.Equal(customResponse.message, response["message"])
	}
}

func TestValidateCustomResponseSuite(t *testing.T) {
	suite.Run(t, &ValidateCustomResponseSuite{})
}
