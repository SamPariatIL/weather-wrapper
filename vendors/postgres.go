package vendors

import (
	"fmt"
	"github.com/SamPariatIL/weather-wrapper/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var postgresDB *gorm.DB

func InitPostgres() {
	cfg := config.GetConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.PostgresConfig.Host,
		cfg.PostgresConfig.User,
		cfg.PostgresConfig.Password,
		cfg.PostgresConfig.Database,
		cfg.PostgresConfig.Port,
		cfg.PostgresConfig.SSLMode,
		cfg.PostgresConfig.TimeZone,
	)

	var err error

	postgresDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
}

func GetPostgresDB() *gorm.DB {
	return postgresDB
}
