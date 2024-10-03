package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SamPariatIL/weather-wrapper/entities"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

type WeatherRepository interface {
	GetCurrentWeather(ctx context.Context, latitude, longitude float32) (*entities.CurrentWeather, error)
	GetFiveDayForecast(ctx context.Context, latitude, longitude float32) (*entities.Forecast, error)
	SetCurrentWeather(ctx context.Context, latitude, longitude float32, currentWeather *entities.CurrentWeather) error
	SetFiveDayForecast(ctx context.Context, latitude, longitude float32, forecast *entities.Forecast) error
}

type weatherRepository struct {
	redisClient *redis.Client
	logger      *zap.Logger
}

func NewWeatherRepository(rc *redis.Client, zl *zap.Logger) WeatherRepository {
	return &weatherRepository{
		redisClient: rc,
		logger:      zl,
	}
}

func (wr *weatherRepository) GetCurrentWeather(ctx context.Context, latitude, longitude float32) (*entities.CurrentWeather, error) {
	key := getCurrentWeatherKey(latitude, longitude)

	weatherJSON, err := wr.redisClient.Get(ctx, key).Result()

	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var currentWeather entities.CurrentWeather

	err = json.Unmarshal([]byte(weatherJSON), &currentWeather)
	if err != nil {
		return nil, err
	}

	wr.logger.Info(fmt.Sprintf("fetched cached weather for %f, %f", latitude, longitude))

	return &currentWeather, nil
}

func (wr *weatherRepository) GetFiveDayForecast(ctx context.Context, latitude, longitude float32) (*entities.Forecast, error) {
	key := getFiveDayWeatherKey(latitude, longitude)

	weatherJSON, err := wr.redisClient.Get(ctx, key).Result()

	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var fiveDayForecast entities.Forecast

	err = json.Unmarshal([]byte(weatherJSON), &fiveDayForecast)
	if err != nil {
		return nil, err
	}

	wr.logger.Info(fmt.Sprintf("fetched cached five day weather for %f, %f", latitude, longitude))

	return &fiveDayForecast, nil
}

func (wr *weatherRepository) SetCurrentWeather(ctx context.Context, latitude, longitude float32, currentWeather *entities.CurrentWeather) error {
	key := getCurrentWeatherKey(latitude, longitude)

	weatherJSON, err := json.Marshal(currentWeather)
	if err != nil {
		return err
	}

	err = wr.redisClient.SetEx(ctx, key, weatherJSON, time.Hour).Err()
	if err != nil {
		return err
	}

	wr.logger.Info(fmt.Sprintf("saved current weather for %f, %f", latitude, longitude))

	return nil
}

func (wr *weatherRepository) SetFiveDayForecast(ctx context.Context, latitude, longitude float32, fiveDayForecast *entities.Forecast) error {
	key := getFiveDayWeatherKey(latitude, longitude)

	weatherJSON, err := json.Marshal(fiveDayForecast)
	if err != nil {
		return err
	}

	err = wr.redisClient.SetEx(ctx, key, weatherJSON, time.Hour).Err()
	if err != nil {
		return err
	}

	wr.logger.Info(fmt.Sprintf("saved five day weather for %f, %f", latitude, longitude))

	return nil
}

func getCurrentWeatherKey(latitude, longitude float32) string {
	return fmt.Sprintf("current_weather_%f_%f", latitude, longitude)
}

func getFiveDayWeatherKey(latitude, longitude float32) string {
	return fmt.Sprintf("five_day_weather_%f_%f", latitude, longitude)
}
