package handlers

import (
	"github.com/SamPariatIL/weather-wrapper/services"
	"github.com/SamPariatIL/weather-wrapper/utils"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type WeatherHandler interface {
	GetCurrentWeather(ctx fiber.Ctx) error
	GetFiveDayForecast(ctx fiber.Ctx) error
}

type weatherHandler struct {
	weatherService services.WeatherService
	logger         *zap.Logger
}

func NewWeatherHandler(ws services.WeatherService, zl *zap.Logger) WeatherHandler {
	return &weatherHandler{
		weatherService: ws,
		logger:         zl,
	}
}

func (wh *weatherHandler) GetCurrentWeather(ctx fiber.Ctx) error {
	lat, lon, err := utils.ValidateLatLon(ctx.Query("lat"), ctx.Query("long"))
	if err != nil {
		wh.logger.Warn(invalidLatLon)
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.CustomResponse(nil, fiber.StatusBadRequest, invalidLatLon, err.Error()))
	}

	currentWeather, err := wh.weatherService.GetCurrentWeather(lat, lon)
	if err != nil {
		wh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.CustomResponse(nil, fiber.StatusInternalServerError, weatherFetchingError, err.Error()))
	}

	wh.logger.Info(successFetchingWeather)
	return ctx.Status(fiber.StatusOK).JSON(utils.CustomResponse(currentWeather, fiber.StatusOK, "", successFetchingWeather))
}

func (wh *weatherHandler) GetFiveDayForecast(ctx fiber.Ctx) error {
	lat, lon, err := utils.ValidateLatLon(ctx.Query("lat"), ctx.Query("long"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.CustomResponse(nil, fiber.StatusBadRequest, invalidLatLon, err.Error()))
	}

	forecast, err := wh.weatherService.GetFiveDayForecast(lat, lon)
	if err != nil {
		wh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.CustomResponse(nil, fiber.StatusInternalServerError, weatherFetchingError, err.Error()))
	}

	wh.logger.Info(successFetchingWeather)
	return ctx.Status(fiber.StatusOK).JSON(utils.CustomResponse(forecast, fiber.StatusOK, "", successFetchingWeather))
}
