package cmd

import (
	"github.com/SamPariatIL/weather-wrapper/config"
	"github.com/SamPariatIL/weather-wrapper/handlers"
	"github.com/SamPariatIL/weather-wrapper/services"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"log"
)

func setupRoutes(app *fiber.App, logger *zap.Logger) {
	weatherService := services.NewWeatherService()
	weatherHandler := handlers.NewWeatherHandler(weatherService, logger)

	geocodingService := services.NewGeocodingService()
	geocodingHandler := handlers.NewGeocodingHandler(geocodingService, logger)

	api := app.Group("/api")
	v1 := api.Group("/v1")

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
	app := fiber.New()

	setupRoutes(app, logger)

	err = app.Listen(":8181")
	if err != nil {
		log.Fatal("An error occurred setting up Fiber", err)
	}
}
