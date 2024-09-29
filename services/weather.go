package services

import (
	"encoding/json"
	"fmt"
	"github.com/SamPariatIL/weather-wrapper/config"
	"github.com/SamPariatIL/weather-wrapper/entities"
	"net/http"
)

type WeatherService interface {
	GetCurrentWeather(latitude, longitude float32) (*entities.CurrentWeather, error)
	GetFiveDayForecast(latitude, longitude float32) (*entities.Forecast, error)
}

type weatherService struct {
}

func NewWeatherService() WeatherService {
	return &weatherService{}
}

func (ws *weatherService) GetCurrentWeather(latitude, longitude float32) (*entities.CurrentWeather, error) {
	conf := config.GetConfig()

	url := fmt.Sprintf(
		"https://%s/weather?lat=%f&lon=%f&units=metric&appid=%s",
		conf.WeatherConfig.BaseURL,
		latitude,
		longitude,
		conf.WeatherConfig.APIKey,
	)

	var err error

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

	return &currentWeather, err
}

func (ws *weatherService) GetFiveDayForecast(latitude, longitude float32) (*entities.Forecast, error) {
	conf := config.GetConfig()

	url := fmt.Sprintf(
		"https://%s/forecast?lat=%f&lon=%f&units=metric&appid=%s",
		conf.WeatherConfig.BaseURL,
		latitude,
		longitude,
		conf.WeatherConfig.APIKey,
	)

	var err error

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

	return &forecast, err
}
