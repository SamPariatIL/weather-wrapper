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

type GeocodingRepository interface {
	GetGeocodeForCity(ctx context.Context, city string, limit int) (*entities.Coord, error)
	GetCityFromLatLon(ctx context.Context, lat, lon float32) (*string, error)
	SetGeocodeForCity(ctx context.Context, city string, limit int, coord *entities.Coord) error
	SetCityFromLatLon(ctx context.Context, lat, lon float32, city string) error
}

type geocodingRepository struct {
	redisClient *redis.Client
	logger      *zap.Logger
}

func NewGeocodingRepository(rc *redis.Client, zl *zap.Logger) GeocodingRepository {
	return &geocodingRepository{
		redisClient: rc,
		logger:      zl,
	}
}

func (gr *geocodingRepository) GetGeocodeForCity(ctx context.Context, city string, limit int) (*entities.Coord, error) {
	key := getGeocodeKey(city, limit)

	coordJSON, err := gr.redisClient.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var coords entities.Coord

	err = json.Unmarshal([]byte(coordJSON), &coords)
	if err != nil {
		return nil, err
	}

	gr.logger.Info(fmt.Sprintf("fetched cached geocode for %s, %d", city, limit))

	return nil, nil
}

func (gr *geocodingRepository) GetCityFromLatLon(ctx context.Context, lat, lon float32) (*string, error) {
	key := getCityKey(lat, lon)

	city, err := gr.redisClient.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	gr.logger.Info(fmt.Sprintf("fetched cached city for %f, %f", lat, lon))

	return &city, nil
}

func (gr *geocodingRepository) SetGeocodeForCity(ctx context.Context, city string, limit int, coord *entities.Coord) error {
	key := getGeocodeKey(city, limit)

	coordJSON, err := json.Marshal(coord)
	if err != nil {
		return err
	}

	err = gr.redisClient.Set(ctx, key, coordJSON, time.Hour*24).Err()
	if err != nil {
		return err
	}

	gr.logger.Info(fmt.Sprintf("saved geocode for %s", city))

	return nil
}

func (gr *geocodingRepository) SetCityFromLatLon(ctx context.Context, lat, lon float32, city string) error {
	key := getCityKey(lat, lon)

	err := gr.redisClient.Set(ctx, key, city, time.Hour*24).Err()
	if err != nil {
		return err
	}

	gr.logger.Info(fmt.Sprintf("saved city for %f, %f", lat, lon))

	return nil
}

func getCityKey(latitude, longitude float32) string {
	return fmt.Sprintf("reverse_geocode_%f_%f", latitude, longitude)
}

func getGeocodeKey(city string, limit int) string {
	return fmt.Sprintf("geocode_%s_%d", city, limit)
}
