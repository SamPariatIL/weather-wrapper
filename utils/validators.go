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

func ValidateLimit(limitString string) (int, error) {
	if limitString == "" {
		return 5, nil
	}

	limit, err := strconv.ParseInt(limitString, 10, 32)
	if err != nil {
		return 0, errors.New("limit is not a number")
	}
	if limit < 1 || limit > 10 {
		return 0, errors.New("limit must be between 1 and 10")
	}

	return int(limit), nil
}

func ValidateCity(city string) error {
	if city == "" {
		return errors.New("city is empty")
	}

	return nil
}

func ValidateDateRange(startDate, endDate string) (int64, int64, error) {
	if startDate == "" || endDate == "" {
		return 0, 0, errors.New("date is empty")
	}

	startDateInt, err := strconv.ParseInt(startDate, 10, 64)
	if err != nil {
		return 0, 0, errors.New("start date is not a number")
	}

	endDateInt, err := strconv.ParseInt(endDate, 10, 64)
	if err != nil {
		return 0, 0, errors.New("end date is not a number")
	}

	if startDateInt > endDateInt {
		return 0, 0, errors.New("start date must be before end date")
	}

	if startDateInt < 0 || endDateInt < 0 {
		return 0, 0, errors.New("date must be positive")
	}

	return startDateInt, endDateInt, nil
}
