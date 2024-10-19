package config

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	FirebaseConfig FirebaseConfig
	GeocodeConfig  GeocodeConfig
	RedisConfig    RedisConfig
	PostgresConfig PostgresConfig
	WeatherConfig  WeatherConfig
}

type GeocodeConfig struct {
	APIKey  string
	BaseURL string
}

type WeatherConfig struct {
	APIKey  string
	BaseURL string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	Timeout  time.Duration
}

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	Database string
	Port     int
	SSLMode  string
	TimeZone string
}

type FirebaseConfig struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x_509_cert_url"`
	ClientX509CertURL       string `json:"client_x_509_cert_url"`
	UniverseDomain          string `json:"universe_domain"`
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

func parseEnvInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		log.Printf("Failed to parse %s, using fallback %d", key, fallback)
	}

	return fallback
}

func loadConfig() (*Config, error) {
	var config Config

	config.FirebaseConfig = FirebaseConfig{
		Type:                    getEnv("FIREBASE_TYPE", ""),
		ProjectID:               getEnv("FIREBASE_PROJECT_ID", ""),
		PrivateKeyID:            getEnv("FIREBASE_PRIVATE_KEY_ID", ""),
		PrivateKey:              getEnv("FIREBASE_PRIVATE_KEY", ""),
		ClientEmail:             getEnv("FIREBASE_CLIENT_EMAIL", ""),
		ClientID:                getEnv("FIREBASE_CLIENT_ID", ""),
		AuthURI:                 "https://accounts.google.com/" + getEnv("FIREBASE_AUTH_URI", ""),
		TokenURI:                "https://oauth2.googleapis.com/" + getEnv("FIREBASE_TOKEN_URI", ""),
		AuthProviderX509CertURL: "https://www.googleapis.com/" + getEnv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL", ""),
		ClientX509CertURL:       "https://www.googleapis.com/" + getEnv("FIREBASE_CLIENT_X509_CERT_URL", ""),
		UniverseDomain:          getEnv("FIREBASE_UNIVERSE_DOMAIN", ""),
	}

	config.GeocodeConfig = GeocodeConfig{
		APIKey:  getEnv("GEOCODE_API_KEY", ""),
		BaseURL: getEnv("GEOCODE_BASE_URL", ""),
	}

	config.WeatherConfig = WeatherConfig{
		APIKey:  getEnv("WEATHER_API_KEY", ""),
		BaseURL: getEnv("WEATHER_BASE_URL", ""),
	}

	config.RedisConfig = RedisConfig{
		Addr:     getEnv("REDIS_ADDRESS", ""),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       parseEnvInt("REDIS_DB", 0),
		Timeout:  10,
	}

	config.PostgresConfig = PostgresConfig{
		Database: getEnv("POSTGRES_DATABASE", ""),
		Password: getEnv("POSTGRES_PASSWORD", ""),
		Host:     getEnv("POSTGRES_HOST", "localhost"),
		Port:     parseEnvInt("POSTGRES_PORT", 5432),
		SSLMode:  getEnv("POSTGRES_SSL_MODE", "disable"),
		User:     getEnv("POSTGRES_USER", ""),
		TimeZone: getEnv("POSTGRES_TIMEZONE", "Asia/Shanghai"),
	}

	return &config, nil
}

func GetConfig() *Config {
	once.Do(func() {
		var err error
		conf, err = loadConfig()
		if err != nil {
			log.Fatalf("Failed to load configuration: %v", err)
		}
	})

	return conf
}
