package services

import (
	"encoding/json"
	"fmt"
	"github.com/SamPariatIL/weather-wrapper/config"
	"github.com/SamPariatIL/weather-wrapper/entities"
	"go.uber.org/zap"
	"net/http"
)

type GeocodingService interface {
	GetGeocodeForCity(city string, limit uint) (*entities.Coord, error)
	GetCityFromLatLon(lat float32, lon float32) (string, error)
}

type geocodingService struct {
	logger *zap.Logger
}

func NewGeocodingService(zl *zap.Logger) GeocodingService {
	return &geocodingService{
		logger: zl,
	}
}

func (gs *geocodingService) GetGeocodeForCity(city string, limit uint) (*entities.Coord, error) {
	conf := config.GetConfig()

	url := fmt.Sprintf(
		"https://%s/direct?q=%s&limit=%d&appid=%s",
		conf.GeocodeConfig.BaseURL,
		city,
		limit,
		conf.GeocodeConfig.APIKey,
	)

	var err error

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

	return &entities.Coord{Lat: geocodes[0].Lat, Lon: geocodes[0].Lon}, nil
}

func (gs *geocodingService) GetCityFromLatLon(lat float32, lon float32) (string, error) {
	conf := config.GetConfig()

	url := fmt.Sprintf(
		"https://%s/reverse?lat=%f&lon=%f&limit=1&appid=%s",
		conf.GeocodeConfig.BaseURL,
		lat,
		lon,
		conf.GeocodeConfig.APIKey,
	)

	var err error

	resp, err := http.Get(url)
	if err != nil {
		return "", nil
	}
	defer func() {
		err = resp.Body.Close()
	}()

	var geocodes []entities.Geocode

	err = json.NewDecoder(resp.Body).Decode(&geocodes)
	if err != nil {
		return "", err
	}

	return geocodes[0].Name, nil
}
