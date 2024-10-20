package tests

import (
	"github.com/SamPariatIL/weather-wrapper/utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ValidateCitySuite struct {
	suite.Suite
}

func (suite *ValidateLimitSuite) TestValidCity() {
	cityPairs := []struct {
		city        string
		expectedErr error
	}{
		{"San Francisco", nil},
		{"New York", nil},
		{"Paris", nil},
		{"London", nil},
		{"Berlin", nil},
		{"Munich", nil},
		{"Tokyo", nil},
		{"Moscow", nil},
		{"Seoul", nil},
		{"Sydney", nil},
		{"New Delhi", nil},
	}

	for _, pair := range cityPairs {
		err := utils.ValidateCity(pair.city)
		suite.Nil(err)
	}
}

func (suite *ValidateLimitSuite) TestInvalidCity() {
	err := utils.ValidateCity("")
	suite.Error(err, "city is required")
}

func TestValidateCitySuite(t *testing.T) {
	suite.Run(t, &ValidateCitySuite{})
}
