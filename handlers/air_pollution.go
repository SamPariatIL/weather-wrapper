package handlers

import (
	"github.com/SamPariatIL/weather-wrapper/services"
	"github.com/SamPariatIL/weather-wrapper/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AirPollutionHandler interface {
	GetCurrentAirPollution(ctx *fiber.Ctx) error
	GetAirPollutionForecast(ctx *fiber.Ctx) error
	GetHistoricalAirPollution(ctx *fiber.Ctx) error
}

type airPollutionHandler struct {
	airPollutionService services.AirPollutionService
	logger              *zap.Logger
}

func NewAirPollutionHandler(as services.AirPollutionService, zl *zap.Logger) AirPollutionHandler {
	return &airPollutionHandler{
		airPollutionService: as,
		logger:              zl,
	}
}

// GetCurrentAirPollution godoc
// @Summary Get current air pollution
// @Description Get current air pollution for a given city
// @Tags air-pollution
// @Accept json
// @Produce json
// @Param lat query string true "Latitude"
// @Param long query string true "Longitude"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /air-pollution/now [get]
func (ah *airPollutionHandler) GetCurrentAirPollution(ctx *fiber.Ctx) error {
	lat, lon, err := utils.ValidateLatLon(ctx.Query("lat"), ctx.Query("long"))
	if err != nil {
		ah.logger.Warn(invalidLatLon)
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.CustomResponse(nil, fiber.StatusBadRequest, invalidLatLon, err.Error()))
	}

	currentAirPollution, err := ah.airPollutionService.GetCurrentAirPollution(lat, lon)
	if err != nil {
		ah.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(utils.CustomResponse(nil, fiber.StatusInternalServerError, airPollutionFetchingError, err.Error()))
	}

	ah.logger.Info(successFetchingAirPollution)
	return ctx.Status(fiber.StatusOK).
		JSON(utils.CustomResponse(currentAirPollution, fiber.StatusOK, "", successFetchingAirPollution))
}

// GetAirPollutionForecast godoc
// @Summary Get air pollution forecast
// @Description Get air pollution forecast for a given city
// @Tags air-pollution
// @Accept json
// @Produce json
// @Param lat query string true "Latitude"
// @Param long query string true "Longitude"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /air-pollution/forecast [get]
func (ah *airPollutionHandler) GetAirPollutionForecast(ctx *fiber.Ctx) error {
	lat, lon, err := utils.ValidateLatLon(ctx.Query("lat"), ctx.Query("long"))
	if err != nil {
		ah.logger.Warn(invalidLatLon)
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.CustomResponse(nil, fiber.StatusBadRequest, invalidLatLon, err.Error()))
	}

	airPollutionForecast, err := ah.airPollutionService.GetAirPollutionForecast(lat, lon)
	if err != nil {
		ah.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(utils.CustomResponse(nil, fiber.StatusInternalServerError, airPollutionFetchingError, err.Error()))
	}

	ah.logger.Info(successFetchingAirPollution)
	return ctx.Status(fiber.StatusOK).
		JSON(utils.CustomResponse(airPollutionForecast, fiber.StatusOK, "", successFetchingAirPollution))
}

// GetHistoricalAirPollution godoc
// @Summary Get historical air pollution
// @Description Get historical air pollution for a given city
// @Tags air-pollution
// @Accept json
// @Produce json
// @Param lat query string true "Latitude"
// @Param long query string true "Longitude"
// @Param start query string true "Start Date (Epoch)"
// @Param end query string true "End Date (Epoch)"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /air-pollution/history [get]
func (ah *airPollutionHandler) GetHistoricalAirPollution(ctx *fiber.Ctx) error {
	lat, lon, err := utils.ValidateLatLon(ctx.Query("lat"), ctx.Query("long"))
	if err != nil {
		ah.logger.Warn(invalidLatLon)
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.CustomResponse(nil, fiber.StatusBadRequest, invalidLatLon, err.Error()))
	}

	startDate, endDate, err := utils.ValidateDateRange(ctx.Query("start"), ctx.Query("end"))
	if err != nil {
		ah.logger.Warn(invalidDate)
		return ctx.Status(fiber.StatusBadRequest).
			JSON(utils.CustomResponse(nil, fiber.StatusBadRequest, invalidDate, err.Error()))
	}

	airPollutionHistory, err := ah.airPollutionService.GetHistoricalAirPollution(lat, lon, startDate, endDate)
	if err != nil {
		ah.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(utils.CustomResponse(nil, fiber.StatusInternalServerError, airPollutionFetchingError, err.Error()))
	}

	ah.logger.Info(successFetchingAirPollution)
	return ctx.Status(fiber.StatusOK).
		JSON(utils.CustomResponse(airPollutionHistory, fiber.StatusOK, "", successFetchingAirPollution))
}
