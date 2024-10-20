package tests

import (
	"github.com/SamPariatIL/weather-wrapper/utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ValidateLatLonSuite struct {
	suite.Suite
}

func (suite *ValidateLatLonSuite) TestValidLatLon() {
	validLatLonPairs := []struct {
		latString string
		lonString string
		lat       float32
		lon       float32
	}{
		{"0", "0", 0.0, 0.0},
		{"40.7128", "-74.0060", 40.7128, -74.0060},
		{"-33.8688", "151.2093", -33.8688, 151.2093},
		{"51.5074", "-0.1278", 51.5074, -0.1278},
		{"-90", "180", -90.0, 180.0},
		{"90", "-180", 90.0, -180.0},
		{"37.7749", "-122.4194", 37.7749, -122.4194},
		{"-45.0", "45.0", -45.0, 45.0},
		{"25.7617", "-80.1918", 25.7617, -80.1918},
	}

	for _, latLon := range validLatLonPairs {
		lat, lon, err := utils.ValidateLatLon(latLon.latString, latLon.lonString)
		suite.Nil(err)
		suite.Equal(latLon.lat, lat)
		suite.Equal(latLon.lon, lon)
	}
}

func (suite *ValidateLatLonSuite) TestInvalidLatLon() {
	invalidLatLonPairs := []struct {
		latString   string
		lonString   string
		expectedErr string
	}{
		{"", "0", "latitude is required"},
		{"0", "", "longitude is required"},
		{"", "", "latitude is required"},
		{"invalid", "0", "latitude is not a number"},
		{"0", "invalid", "longitude is required"},
		{"invalid", "invalid", "latitude is not a number"},
		{"100", "0", "latitude must be between -90 and 90"},
		{"-100", "0", "latitude must be between -90 and 90"},
		{"0", "200", "longitude must be between -180 and 180"},
		{"0", "-200", "longitude must be between -180 and 180"},
	}

	for _, latLon := range invalidLatLonPairs {
		lat, lon, err := utils.ValidateLatLon(latLon.latString, latLon.lonString)
		suite.Error(err, latLon.expectedErr)
		suite.Equal(lat, float32(0.0))
		suite.Equal(lon, float32(0.0))
	}
}

func TestValidateLatLonSuite(t *testing.T) {
	suite.Run(t, &ValidateLatLonSuite{})
}
