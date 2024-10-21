package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SamPariatIL/weather-wrapper/config"
	"github.com/SamPariatIL/weather-wrapper/entities"
	"github.com/SamPariatIL/weather-wrapper/repository"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type AirPollutionService interface {
	GetCurrentAirPollution(latitude, longitude float32) (*entities.AirPollution, error)
	GetAirPollutionForecast(latitude, longitude float32) (*entities.AirPollution, error)
	GetHistoricalAirPollution(latitude, longitude float32, start, end int64) (*entities.AirPollution, error)
}

type airPollutionService struct {
	airPollutionRepo repository.AirPollutionRepository
	logger           *zap.Logger
}

func NewAirPollutionService(ar repository.AirPollutionRepository, zl *zap.Logger) AirPollutionService {
	return &airPollutionService{
		airPollutionRepo: ar,
		logger:           zl,
	}
}

func (as *airPollutionService) GetCurrentAirPollution(latitude, longitude float32) (*entities.AirPollution, error) {
	conf := config.GetConfig()

	var err error

	savedAirPollution, err := as.airPollutionRepo.GetCurrentAirPollution(context.Background(), latitude, longitude)
	if err != nil {
		return nil, err
	}

	if savedAirPollution != nil {
		return savedAirPollution, nil
	}

	url := fmt.Sprintf(
		"https://%s?lat=%f&lon=%f&appid=%s",
		conf.AirPollutionConfig.BaseURL,
		latitude,
		longitude,
		conf.AirPollutionConfig.APIKey,
	)

	log.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, nil
	}

	defer func() {
		err = resp.Body.Close()
	}()

	var airPollution entities.AirPollution

	err = json.NewDecoder(resp.Body).Decode(&airPollution)
	if err != nil {
		return nil, err
	}

	err = as.airPollutionRepo.SetCurrentAirPollution(context.Background(), latitude, longitude, &airPollution)
	if err != nil {
		return nil, err
	}

	return &airPollution, nil
}

func (as *airPollutionService) GetAirPollutionForecast(latitude, longitude float32) (*entities.AirPollution, error) {
	conf := config.GetConfig()

	var err error

	savedAirPollutionForecast, err := as.airPollutionRepo.GetAirPollutionForecast(context.Background(), latitude, longitude)
	if err != nil {
		return nil, err
	}

	if savedAirPollutionForecast != nil {
		return savedAirPollutionForecast, nil
	}

	url := fmt.Sprintf(
		"https://%s/forecast?lat=%f&lon=%f&appid=%s",
		conf.AirPollutionConfig.BaseURL,
		latitude,
		longitude,
		conf.AirPollutionConfig.APIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, nil
	}

	defer func() {
		err = resp.Body.Close()
	}()

	var airPollutionForecast entities.AirPollution

	err = json.NewDecoder(resp.Body).Decode(&airPollutionForecast)
	if err != nil {
		return nil, err
	}

	err = as.airPollutionRepo.SetAirPollutionForecast(context.Background(), latitude, longitude, &airPollutionForecast)
	if err != nil {
		return nil, err
	}

	return &airPollutionForecast, nil
}

func (as *airPollutionService) GetHistoricalAirPollution(latitude, longitude float32, start, end int64) (*entities.AirPollution, error) {
	conf := config.GetConfig()

	var err error

	savedHistoricalAirPollution, err := as.airPollutionRepo.GetHistoricalAirPollution(context.Background(), latitude, longitude)
	if err != nil {
		return nil, err
	}

	if savedHistoricalAirPollution != nil {
		return savedHistoricalAirPollution, nil
	}

	url := fmt.Sprintf(
		"https://%s/history?lat=%f&lon=%f&start=%d&end=%d&appid=%s",
		conf.AirPollutionConfig.BaseURL,
		latitude,
		longitude,
		start,
		end,
		conf.AirPollutionConfig.APIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, nil
	}

	defer func() {
		err = resp.Body.Close()
	}()

	var historicalAirPollution entities.AirPollution

	err = json.NewDecoder(resp.Body).Decode(&historicalAirPollution)
	if err != nil {
		return nil, err
	}

	err = as.airPollutionRepo.SetHistoricalAirPollution(context.Background(), latitude, longitude, &historicalAirPollution)
	if err != nil {
		return nil, err
	}

	return &historicalAirPollution, nil
}
