package handlers

import (
	"github.com/SamPariatIL/weather-wrapper/services"
	"github.com/SamPariatIL/weather-wrapper/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type GeocodingHandler interface {
	GetGeocodeForCity(ctx *fiber.Ctx) error
	GetCityFromLatLon(ctx *fiber.Ctx) error
}

type geocodingHandler struct {
	geocodingService services.GeocodingService
	logger           *zap.Logger
}

func NewGeocodingHandler(gs services.GeocodingService, zl *zap.Logger) GeocodingHandler {
	return &geocodingHandler{
		geocodingService: gs,
		logger:           zl,
	}
}

// GetGeocodeForCity godoc
// @Summary Get geocoding
// @Description Get geocoding for a given city
// @Tags geocode
// @Accept json
// @Produce json
// @Param city query string true "City"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /geocode [get]
func (gh *geocodingHandler) GetGeocodeForCity(ctx *fiber.Ctx) error {
	city := ctx.Query("city")
	err := utils.ValidateCity(city)
	if err != nil {
		gh.logger.Warn(invalidCity)
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.CustomResponse(nil, fiber.StatusBadRequest, invalidCity, err.Error()))
	}

	limit, err := utils.ValidateLimit(ctx.Query("limit"))
	if err != nil {
		gh.logger.Warn(invalidLimit)
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.CustomResponse(nil, fiber.StatusBadRequest, invalidLimit, err.Error()))
	}

	coords, err := gh.geocodingService.GetGeocodeForCity(city, limit)
	if err != nil {
		gh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.CustomResponse(nil, fiber.StatusInternalServerError, geocodingFetchingError, err.Error()))
	}

	gh.logger.Info(successFetchingGeocode)
	return ctx.Status(fiber.StatusOK).JSON(utils.CustomResponse(coords, fiber.StatusOK, "", successFetchingGeocode))
}

// GetCityFromLatLon godoc
// @Summary Get city
// @Description Get city from latitude and longitude
// @Tags geocode
// @Accept json
// @Produce json
// @Param lat query string true "Latitude"
// @Param long query string true "Longitude"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /geocode/reverse [get]
func (gh *geocodingHandler) GetCityFromLatLon(ctx *fiber.Ctx) error {
	lat, lon, err := utils.ValidateLatLon(ctx.Query("lat"), ctx.Query("long"))
	if err != nil {
		gh.logger.Warn(invalidLatLon)
		return ctx.Status(fiber.StatusBadRequest).JSON(utils.CustomResponse(nil, fiber.StatusBadRequest, invalidLatLon, err.Error()))
	}

	city, err := gh.geocodingService.GetCityFromLatLon(lat, lon)

	if city == "" && err != nil {
		gh.logger.Warn(invalidLatLon)
		return ctx.Status(fiber.StatusNotFound).JSON(utils.CustomResponse(nil, fiber.StatusNotFound, invalidLatLon, err.Error()))
	}

	if err != nil {
		gh.logger.Error(err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.CustomResponse(nil, fiber.StatusInternalServerError, reverseGeocodingFetchingError, err.Error()))
	}

	gh.logger.Info(successFetchingReverseGeocoding)
	return ctx.Status(fiber.StatusOK).JSON(utils.CustomResponse(city, fiber.StatusOK, "", successFetchingReverseGeocoding))
}
