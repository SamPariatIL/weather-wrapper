package utils

import (
	"errors"
	"strconv"
)

func ValidateLatLon(latString, lonString string) (float32, float32, error) {
	if latString == "" {
		return 0.0, 0.0, errors.New("latitude is required")
	}
	if lonString == "" {
		return 0.0, 0.0, errors.New("longitude is required")
	}

	lat, err := strconv.ParseFloat(latString, 32)
	if err != nil {
		return 0.0, 0.0, errors.New("latitude is not a number")
	}

	lon, err := strconv.ParseFloat(lonString, 32)
	if err != nil {
		return 0.0, 0.0, errors.New("longitude is required")
	}

	if lat < -90.0 || lat > 90.0 {
		return 0.0, 0.0, errors.New("latitude must be between -90 and 90")
	}
	if lon < -180.0 || lon > 180.0 {
		return 0.0, 0.0, errors.New("longitude must be between -180 and 180")
	}

	return float32(lat), float32(lon), nil
}

func ValidateLimit(limitString string) (uint, error) {
	if limitString == "" {
		return 5, nil
	}

	limit, err := strconv.ParseUint(limitString, 10, 32)
	if err != nil {
		return 0, errors.New("limit is not a number")
	}

	return uint(limit), nil
}

func ValidateCity(city string) error {
	if city == "" {
		return errors.New("city is empty")
	}

	return nil
}
