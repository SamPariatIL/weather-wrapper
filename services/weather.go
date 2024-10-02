package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SamPariatIL/weather-wrapper/config"
	"github.com/SamPariatIL/weather-wrapper/entities"
	"github.com/SamPariatIL/weather-wrapper/repository"
	"go.uber.org/zap"
	"net/http"
)

type WeatherService interface {
	GetCurrentWeather(latitude, longitude float32) (*entities.CurrentWeather, error)
	GetFiveDayForecast(latitude, longitude float32) (*entities.Forecast, error)
}

type weatherService struct {
	weatherRepo repository.WeatherRepository
	logger      *zap.Logger
}

func NewWeatherService(wr repository.WeatherRepository, zl *zap.Logger) WeatherService {
	return &weatherService{
		weatherRepo: wr,
		logger:      zl,
	}
}

func (ws *weatherService) GetCurrentWeather(latitude, longitude float32) (*entities.CurrentWeather, error) {
	conf := config.GetConfig()

	var err error

	savedWeather, err := ws.weatherRepo.GetCurrentWeather(context.Background(), latitude, longitude)
	if err != nil {
		return nil, err
	}

	if savedWeather != nil {
		return savedWeather, nil
	}

	url := fmt.Sprintf(
		"https://%s/weather?lat=%f&lon=%f&units=metric&appid=%s",
		conf.WeatherConfig.BaseURL,
		latitude,
		longitude,
		conf.WeatherConfig.APIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = resp.Body.Close()
	}()

	var currentWeather entities.CurrentWeather

	err = json.NewDecoder(resp.Body).Decode(&currentWeather)
	if err != nil {
		return nil, err
	}

	err = ws.weatherRepo.SetCurrentWeather(context.Background(), latitude, longitude, &currentWeather)
	if err != nil {
		return nil, err
	}

	return &currentWeather, err
}

func (ws *weatherService) GetFiveDayForecast(latitude, longitude float32) (*entities.Forecast, error) {
	conf := config.GetConfig()

	var err error

	savedForecast, err := ws.weatherRepo.GetFiveDayForecast(context.Background(), latitude, longitude)
	if err != nil {
		return nil, err
	}

	if savedForecast != nil {
		return savedForecast, nil
	}

	url := fmt.Sprintf(
		"https://%s/forecast?lat=%f&lon=%f&units=metric&appid=%s",
		conf.WeatherConfig.BaseURL,
		latitude,
		longitude,
		conf.WeatherConfig.APIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = resp.Body.Close()
	}()

	var forecast entities.Forecast

	err = json.NewDecoder(resp.Body).Decode(&forecast)
	if err != nil {
		return nil, err
	}

	err = ws.weatherRepo.SetFiveDayForecast(context.Background(), latitude, longitude, &forecast)
	if err != nil {
		return nil, err
	}

	return &forecast, err
}
