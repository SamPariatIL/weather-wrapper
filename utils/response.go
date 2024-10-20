package utils

import "github.com/gofiber/fiber/v2"

func CustomResponse(data any, status int, error, message string) fiber.Map {
	return fiber.Map{
		"data":    data,
		"status":  status,
		"error":   error,
		"message": message,
	}
}
