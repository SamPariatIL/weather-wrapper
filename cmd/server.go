package cmd

import (
	"github.com/SamPariatIL/weather-wrapper/config"
	_ "github.com/SamPariatIL/weather-wrapper/docs"
	"github.com/SamPariatIL/weather-wrapper/handlers"
	"github.com/SamPariatIL/weather-wrapper/repository"
	"github.com/SamPariatIL/weather-wrapper/services"
	"github.com/SamPariatIL/weather-wrapper/vendors"
	"github.com/gofiber/fiber/v2"
	fiberCors "github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	fiberRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"go.uber.org/zap"
	"log"
)

func setupRoutes(app *fiber.App, logger *zap.Logger) {
	redisClient := vendors.GetRedisClient()
	authClient := vendors.GetFirebaseAuth()

	userRepo := repository.NewUserRepository(authClient, logger)
	userService := services.NewUserService(userRepo, logger)
	userHandler := handlers.NewUserHandler(userService, logger)

	weatherRepo := repository.NewWeatherRepository(redisClient, logger)
	weatherService := services.NewWeatherService(weatherRepo, logger)
	weatherHandler := handlers.NewWeatherHandler(weatherService, logger)

	geocodingRepo := repository.NewGeocodingRepository(redisClient, logger)
	geocodingService := services.NewGeocodingService(geocodingRepo, logger)
	geocodingHandler := handlers.NewGeocodingHandler(geocodingService, logger)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	health := v1.Group("/")
	health.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Weather wrapper is running woohoo!!")
	})

	apiDocs := v1.Group("/swagger")
	apiDocs.Get("*", swagger.HandlerDefault)

	weatherV1 := v1.Group("/weather")
	weatherV1.Get("/now", weatherHandler.GetCurrentWeather)
	weatherV1.Get("/forecast", weatherHandler.GetFiveDayForecast)

	geocodingV1 := v1.Group("/geocode")
	geocodingV1.Get("/", geocodingHandler.GetGeocodeForCity)
	geocodingV1.Get("/reverse", geocodingHandler.GetCityFromLatLon)

	usersV1 := v1.Group("/users")
	usersV1.Get("/token", userHandler.GenerateToken)
	usersV1.Post("/signup", userHandler.CreateUser)
	usersV1.Post("/verify", userHandler.SendVerificationEmail)
	usersV1.Post("/reset-password", userHandler.ResetPassword)
	usersV1.Put("/:uid", userHandler.UpdateUser)
	usersV1.Delete("/:uid", userHandler.DeleteUser)
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
	app.Use(fiberLogger.New())
	app.Use(fiberRecover.New())
	app.Use(fiberCors.New())

	setupRoutes(app, logger)

	err = app.Listen(":8181")
	if err != nil {
		log.Fatal("An error occurred setting up Fiber", err)
	}
}
