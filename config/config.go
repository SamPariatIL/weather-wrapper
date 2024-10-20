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
		Type:                    getEnv(FirebaseType, ""),
		ProjectID:               getEnv(FirebaseProjectId, ""),
		PrivateKeyID:            getEnv(FirebasePrivateKeyId, ""),
		PrivateKey:              getEnv(FirebasePrivateKey, ""),
		ClientEmail:             getEnv(FirebaseClientEmail, ""),
		ClientID:                getEnv(FirebaseClientId, ""),
		AuthURI:                 "https://accounts.google.com/" + getEnv(FirebaseAuthUri, ""),
		TokenURI:                "https://oauth2.googleapis.com/" + getEnv(FirebaseTokenUri, ""),
		AuthProviderX509CertURL: "https://www.googleapis.com/" + getEnv(FirebaseAuthProviderX509CertUrl, ""),
		ClientX509CertURL:       "https://www.googleapis.com/" + getEnv(FirebaseClientX509CertUrl, ""),
		UniverseDomain:          getEnv(FirebaseUniverseDomain, ""),
	}

	config.GeocodeConfig = GeocodeConfig{
		APIKey:  getEnv(GeocodeApiKey, ""),
		BaseURL: getEnv(GeocodeBaseUrl, ""),
	}

	config.WeatherConfig = WeatherConfig{
		APIKey:  getEnv(WeatherApiKey, ""),
		BaseURL: getEnv(WeatherBaseUrl, ""),
	}

	config.RedisConfig = RedisConfig{
		Addr:     getEnv(RedisAddress, ""),
		Password: getEnv(RedisPassword, ""),
		DB:       parseEnvInt(RedisDB, 0),
		Timeout:  time.Second * time.Duration(parseEnvInt(RedisTimeout, 10)),
	}

	config.PostgresConfig = PostgresConfig{
		Database: getEnv(PostgresDatabase, ""),
		Password: getEnv(PostgresPassword, ""),
		Host:     getEnv(PostgresHost, "localhost"),
		Port:     parseEnvInt(PostgresPort, 5432),
		SSLMode:  getEnv(PostgresSslMode, "disable"),
		User:     getEnv(PostgresUser, ""),
		TimeZone: getEnv(PostgresTimezone, "Asia/Shanghai"),
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
