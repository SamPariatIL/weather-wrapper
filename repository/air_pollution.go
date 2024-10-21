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

type AirPollutionRepository interface {
	GetCurrentAirPollution(ctx context.Context, latitude, longitude float32) (*entities.AirPollution, error)
	GetAirPollutionForecast(ctx context.Context, latitude, longitude float32) (*entities.AirPollution, error)
	GetHistoricalAirPollution(ctx context.Context, latitude, longitude float32) (*entities.AirPollution, error)
	SetCurrentAirPollution(ctx context.Context, latitude, longitude float32, airPollution *entities.AirPollution) error
	SetAirPollutionForecast(ctx context.Context, latitude, longitude float32, airPollutionForecast *entities.AirPollution) error
	SetHistoricalAirPollution(ctx context.Context, latitude, longitude float32, historicalAirPollution *entities.AirPollution) error
}

type airPollutionRepository struct {
	redisClient *redis.Client
	logger      *zap.Logger
}

func NewAirPollutionRepository(rc *redis.Client, zl *zap.Logger) AirPollutionRepository {
	return &airPollutionRepository{
		redisClient: rc,
		logger:      zl,
	}
}

func (ar *airPollutionRepository) GetCurrentAirPollution(ctx context.Context, latitude, longitude float32) (*entities.AirPollution, error) {
	key := getCurrentAirPollutionKey(latitude, longitude)

	airPollutionJSON, err := ar.redisClient.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var airPoll entities.AirPollution

	err = json.Unmarshal([]byte(airPollutionJSON), &airPoll)
	if err != nil {
		return nil, err
	}

	ar.logger.Info(fmt.Sprintf("fetched cached air pollution for %f, %f", latitude, longitude))
	return &airPoll, nil
}

func (ar *airPollutionRepository) GetAirPollutionForecast(ctx context.Context, latitude, longitude float32) (*entities.AirPollution, error) {
	key := getAirPollutionForecastKey(latitude, longitude)

	airPollutionForecastJSON, err := ar.redisClient.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var airPollForecast entities.AirPollution

	err = json.Unmarshal([]byte(airPollutionForecastJSON), &airPollForecast)
	if err != nil {
		return nil, err
	}

	ar.logger.Info(fmt.Sprintf("fetched cached air pollution forecast for %f, %f", latitude, longitude))
	return &airPollForecast, nil
}

func (ar *airPollutionRepository) GetHistoricalAirPollution(ctx context.Context, latitude, longitude float32) (*entities.AirPollution, error) {
	key := getHistoricalAirPollutionKey(latitude, longitude)

	historicalAirPollutionJSON, err := ar.redisClient.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var historicalAirPollution entities.AirPollution

	err = json.Unmarshal([]byte(historicalAirPollutionJSON), &historicalAirPollution)
	if err != nil {
		return nil, err
	}

	ar.logger.Info(fmt.Sprintf("fetched cached historical air pollution for %f, %f", latitude, longitude))
	return &historicalAirPollution, nil
}

func (ar *airPollutionRepository) SetCurrentAirPollution(ctx context.Context, latitude, longitude float32, airPollution *entities.AirPollution) error {
	key := getCurrentAirPollutionKey(latitude, longitude)

	airPollutionJSON, err := json.Marshal(airPollution)
	if err != nil {
		return err
	}

	err = ar.redisClient.Set(ctx, key, airPollutionJSON, time.Minute*5).Err()
	if err != nil {
		return err
	}

	ar.logger.Info(fmt.Sprintf("saved current air pollution for %f, %f", latitude, longitude))
	return nil
}

func (ar *airPollutionRepository) SetAirPollutionForecast(ctx context.Context, latitude, longitude float32, airPollutionForecast *entities.AirPollution) error {
	key := getAirPollutionForecastKey(latitude, longitude)

	airPollutionForecastJSON, err := json.Marshal(airPollutionForecast)
	if err != nil {
		return err
	}

	err = ar.redisClient.Set(ctx, key, airPollutionForecastJSON, time.Minute*5).Err()
	if err != nil {
		return err
	}

	ar.logger.Info(fmt.Sprintf("saved air pollution forecast for %f, %f", latitude, longitude))
	return nil
}

func (ar *airPollutionRepository) SetHistoricalAirPollution(ctx context.Context, latitude, longitude float32, historicalAirPollution *entities.AirPollution) error {
	key := getHistoricalAirPollutionKey(latitude, longitude)

	historicalAirPollutionJSON, err := json.Marshal(historicalAirPollution)
	if err != nil {
		return err
	}

	err = ar.redisClient.Set(ctx, key, historicalAirPollutionJSON, time.Minute*5).Err()
	if err != nil {
		return err
	}

	ar.logger.Info(fmt.Sprintf("saved historical air pollution for %f, %f", latitude, longitude))
	return nil
}

func getCurrentAirPollutionKey(lat, lon float32) string {
	return fmt.Sprintf("current_air_pollution_%f_%f", lat, lon)
}

func getAirPollutionForecastKey(lat, lon float32) string {
	return fmt.Sprintf("air_pollution_forecast_%f_%f", lat, lon)
}

func getHistoricalAirPollutionKey(lat, lon float32) string {
	return fmt.Sprintf("historical_air_pollution_%f_%f", lat, lon)
}
