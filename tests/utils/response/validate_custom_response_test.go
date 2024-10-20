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
	customResponses := []struct {
		data    any
		status  int
		error   string
		message string
	}{
		{data: "ok", status: 200, error: "", message: "ok"},
		{data: "error", status: 201, error: "", message: "created"},
		{data: "error", status: 400, error: "something went wrong", message: "bad request"},
		{data: "error", status: 401, error: "something went wrong", message: "unauthorized"},
		{data: "error", status: 403, error: "something went wrong", message: "forbidden"},
		{data: "error", status: 404, error: "something went wrong", message: "not found"},
		{data: "error", status: 409, error: "something went wrong", message: "conflict"},
		{data: "error", status: 500, error: "something went wrong", message: "internal server error"},
		{data: fiber.Map{"message": "ok"}, status: 200, error: "", message: "ok"},
		{data: fiber.Map{"message": "error"}, status: 201, error: "", message: "created"},
		{data: fiber.Map{"message": "error"}, status: 400, error: "something went wrong", message: "bad request"},
		{data: fiber.Map{"message": "error"}, status: 401, error: "something went wrong", message: "unauthorized"},
		{data: fiber.Map{"message": "error"}, status: 403, error: "something went wrong", message: "forbidden"},
		{data: fiber.Map{"message": "error"}, status: 404, error: "something went wrong", message: "not found"},
		{data: fiber.Map{"message": "error"}, status: 409, error: "something went wrong", message: "conflict"},
		{data: fiber.Map{"message": "error"}, status: 500, error: "something went wrong", message: "internal server error"},
		{data: nil, status: 200, error: "", message: "ok"},
		{data: nil, status: 201, error: "", message: "created"},
		{data: nil, status: 400, error: "something went wrong", message: "bad request"},
		{data: nil, status: 401, error: "something went wrong", message: "unauthorized"},
		{data: nil, status: 403, error: "something went wrong", message: "forbidden"},
		{data: nil, status: 404, error: "something went wrong", message: "not found"},
		{data: nil, status: 409, error: "something went wrong", message: "conflict"},
		{data: nil, status: 500, error: "something went wrong", message: "internal server error"},
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
