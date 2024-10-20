package tests

import (
	"github.com/SamPariatIL/weather-wrapper/utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ValidateLimitSuite struct {
	suite.Suite
}

func (suite *ValidateLimitSuite) TestValidLimit() {
	validLimitPairs := []struct {
		limitString string
		limit       int
	}{
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
		{"10", 10},
	}

	for _, pair := range validLimitPairs {
		limit, err := utils.ValidateLimit(pair.limitString)
		suite.Nil(err)
		suite.Equal(pair.limit, limit)
	}
}

func (suite *ValidateLimitSuite) TestInvalidLimit() {
	limitIsNotANumberError := "limit is not a number"
	limitMustBeBetween1And10Error := "limit must be between 1 and 10"

	invalidLimitPairs := []struct {
		limitString   string
		limit         int
		expectedError *string
	}{
		{"", 5, nil},
		{"invalid", 0, &limitIsNotANumberError},
		{"-1", 0, &limitMustBeBetween1And10Error},
		{"11", 0, &limitMustBeBetween1And10Error},
	}

	for _, pair := range invalidLimitPairs {
		limit, err := utils.ValidateLimit(pair.limitString)

		if pair.expectedError == nil {
			suite.Nil(err)
		} else {
			suite.NotNil(err)
			suite.Equal(*pair.expectedError, err.Error())
		}

		suite.Equal(pair.limit, limit)
	}
}

func TestValidateLimitSuite(t *testing.T) {
	suite.Run(t, &ValidateLimitSuite{})
}
