package tests

import (
	"github.com/SamPariatIL/weather-wrapper/config"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"testing"
	"time"
)

var envMap map[string]string

type ConfigSuite struct {
	suite.Suite
}

func (suite *ConfigSuite) SetupTest() {
	log.Println("[config_test] Setting up test environment...")

	envMap = make(map[string]string)

	envMap[config.FirebaseType] = "firebase_type"
	envMap[config.FirebaseProjectId] = "firebase_project_id"
	envMap[config.FirebasePrivateKeyId] = "firebase_private_key_id"
	envMap[config.FirebasePrivateKey] = "firebase_private_key"
	envMap[config.FirebaseClientEmail] = "firebase_client_email"
	envMap[config.FirebaseClientId] = "firebase_client_id"
	envMap[config.FirebaseAuthUri] = "firebase_auth_uri"
	envMap[config.FirebaseTokenUri] = "firebase_token_uri"
	envMap[config.FirebaseAuthProviderX509CertUrl] = "firebase_auth_provider_x509_cert_url"
	envMap[config.FirebaseClientX509CertUrl] = "firebase_client_x509_cert_url"
	envMap[config.FirebaseUniverseDomain] = "firebase_universe_domain"

	envMap[config.GeocodeApiKey] = "geocode_api_key"
	envMap[config.GeocodeBaseUrl] = "geocode_base_url"

	envMap[config.WeatherApiKey] = "weather_api_key"
	envMap[config.WeatherBaseUrl] = "weather_base_url"

	envMap[config.RedisAddress] = "redis_address"
	envMap[config.RedisPassword] = "redis_password"
	envMap[config.RedisDB] = "0"
	envMap[config.RedisTimeout] = "10"

	envMap[config.PostgresDatabase] = "postgres_database"
	envMap[config.PostgresPassword] = "postgres_password"
	envMap[config.PostgresHost] = "postgres_host"
	envMap[config.PostgresPort] = "5432"
	envMap[config.PostgresSslMode] = "postgres_ssl_mode"
	envMap[config.PostgresTimezone] = "postgres_time_zone"

	for key, value := range envMap {
		err := os.Setenv(key, value)
		if err != nil {
			panic(err)
		}
	}
}

func (suite *ConfigSuite) TearDownSuite() {
	log.Println("[config_test] Tearing down test environment...")

	for _, key := range envMap {
		err := os.Unsetenv(key)
		if err != nil {
			panic(err)
		}
	}
}

func (suite *ConfigSuite) TestGetConfigSuccess() {
	conf := config.GetConfig()

	suite.Equal("firebase_type", conf.FirebaseConfig.Type)
	suite.Equal("firebase_project_id", conf.FirebaseConfig.ProjectID)
	suite.Equal("firebase_private_key_id", conf.FirebaseConfig.PrivateKeyID)
	suite.Equal("firebase_private_key", conf.FirebaseConfig.PrivateKey)
	suite.Equal("firebase_client_email", conf.FirebaseConfig.ClientEmail)
	suite.Equal("firebase_client_id", conf.FirebaseConfig.ClientID)
	suite.Equal("https://accounts.google.com/firebase_auth_uri", conf.FirebaseConfig.AuthURI)
	suite.Equal("https://oauth2.googleapis.com/firebase_token_uri", conf.FirebaseConfig.TokenURI)
	suite.Equal("https://www.googleapis.com/firebase_auth_provider_x509_cert_url", conf.FirebaseConfig.AuthProviderX509CertURL)
	suite.Equal("https://www.googleapis.com/firebase_client_x509_cert_url", conf.FirebaseConfig.ClientX509CertURL)
	suite.Equal("firebase_universe_domain", conf.FirebaseConfig.UniverseDomain)
	suite.Equal("geocode_api_key", conf.GeocodeConfig.APIKey)
	suite.Equal("geocode_base_url", conf.GeocodeConfig.BaseURL)
	suite.Equal("weather_api_key", conf.WeatherConfig.APIKey)
	suite.Equal("weather_base_url", conf.WeatherConfig.BaseURL)
	suite.Equal("redis_address", conf.RedisConfig.Addr)
	suite.Equal("redis_password", conf.RedisConfig.Password)
	suite.Equal(0, conf.RedisConfig.DB)
	suite.Equal(10*time.Second, conf.RedisConfig.Timeout)
	suite.Equal("postgres_database", conf.PostgresConfig.Database)
	suite.Equal("postgres_password", conf.PostgresConfig.Password)
	suite.Equal("postgres_host", conf.PostgresConfig.Host)
	suite.Equal(5432, conf.PostgresConfig.Port)
	suite.Equal("postgres_ssl_mode", conf.PostgresConfig.SSLMode)
	suite.Equal("postgres_time_zone", conf.PostgresConfig.TimeZone)
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, &ConfigSuite{})
}
