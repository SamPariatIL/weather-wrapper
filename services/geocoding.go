package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SamPariatIL/weather-wrapper/config"
	"github.com/SamPariatIL/weather-wrapper/entities"
	"github.com/SamPariatIL/weather-wrapper/repository"
	"go.uber.org/zap"
	"net/http"
)

type GeocodingService interface {
	GetGeocodeForCity(city string, limit int) (*entities.Coord, error)
	GetCityFromLatLon(lat, lon float32) (*string, error)
}

type geocodingService struct {
	geocodingRepo repository.GeocodingRepository
	logger        *zap.Logger
}

func NewGeocodingService(gr repository.GeocodingRepository, zl *zap.Logger) GeocodingService {
	return &geocodingService{
		geocodingRepo: gr,
		logger:        zl,
	}
}

func (gs *geocodingService) GetGeocodeForCity(city string, limit int) (*entities.Coord, error) {
	conf := config.GetConfig()

	var err error

	savedGeocode, err := gs.geocodingRepo.GetGeocodeForCity(context.Background(), city, limit)
	if err != nil {
		return nil, err
	}

	if savedGeocode != nil {
		return savedGeocode, nil
	}

	url := fmt.Sprintf(
		"https://%s/direct?q=%s&limit=%d&appid=%s",
		conf.GeocodeConfig.BaseURL,
		city,
		limit,
		conf.GeocodeConfig.APIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, nil
	}

	defer func() {
		err = resp.Body.Close()
	}()

	var geocodes []entities.Geocode

	err = json.NewDecoder(resp.Body).Decode(&geocodes)
	if err != nil {
		return nil, err
	}

	coords := entities.Coord{
		Lat: geocodes[0].Lat,
		Lon: geocodes[0].Lon,
	}

	err = gs.geocodingRepo.SetGeocodeForCity(context.Background(), city, limit, &coords)
	if err != nil {
		return nil, err
	}

	return &coords, nil
}

func (gs *geocodingService) GetCityFromLatLon(lat, lon float32) (*string, error) {
	conf := config.GetConfig()

	var err error

	savedCity, err := gs.geocodingRepo.GetCityFromLatLon(context.Background(), lat, lon)
	if err != nil {
		return nil, err
	}

	if savedCity != nil {
		return savedCity, nil
	}

	url := fmt.Sprintf(
		"https://%s/reverse?lat=%f&lon=%f&limit=1&appid=%s",
		conf.GeocodeConfig.BaseURL,
		lat,
		lon,
		conf.GeocodeConfig.APIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = resp.Body.Close()
	}()

	var geocodes []entities.Geocode

	err = json.NewDecoder(resp.Body).Decode(&geocodes)
	if err != nil {
		return nil, err
	}

	if len(geocodes) == 0 {
		return nil, errors.New("no city found")
	}

	city := geocodes[0].Name

	err = gs.geocodingRepo.SetCityFromLatLon(context.Background(), lat, lon, city)
	if err != nil {
		return nil, err
	}

	return &city, nil
}
