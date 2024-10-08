package cmd

import (
	"github.com/SamPariatIL/weather-wrapper/config"
	"github.com/SamPariatIL/weather-wrapper/handlers"
	"github.com/SamPariatIL/weather-wrapper/repository"
	"github.com/SamPariatIL/weather-wrapper/services"
	"github.com/SamPariatIL/weather-wrapper/vendors"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"log"
)

func setupRoutes(app *fiber.App, logger *zap.Logger) {
	redisClient := vendors.GetRedisClient()

	weatherRepo := repository.NewWeatherRepository(redisClient, logger)
	weatherService := services.NewWeatherService(weatherRepo, logger)
	weatherHandler := handlers.NewWeatherHandler(weatherService, logger)

	geocodingRepo := repository.NewGeocodingRepository(redisClient, logger)
	geocodingService := services.NewGeocodingService(geocodingRepo, logger)
	geocodingHandler := handlers.NewGeocodingHandler(geocodingService, logger)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	health := v1.Group("/")
	health.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Weather wrapper is running woohoo!!")
	})

	weatherV1 := v1.Group("/weather")
	weatherV1.Get("/now", weatherHandler.GetCurrentWeather)
	weatherV1.Get("/forecast", weatherHandler.GetFiveDayForecast)

	geocodingV1 := v1.Group("/geocode")
	geocodingV1.Get("/", geocodingHandler.GetGeocodeForCity)
	geocodingV1.Get("/reverse", geocodingHandler.GetCityFromLatLon)
}

func RunServer() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("An error occurred setting up Zap", err)
	}

	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatal("An error occurred syncing Zap", err)
		}
	}(logger)

	config.GetConfig()
	initVendors()

	app := fiber.New()

	setupRoutes(app, logger)

	err = app.Listen(":8181")
	if err != nil {
		log.Fatal("An error occurred setting up Fiber", err)
	}
}
