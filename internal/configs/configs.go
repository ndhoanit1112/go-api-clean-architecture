package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/ndhoanit1112/go-api-clean-architecture/internal/infrastructure/db"
	"github.com/ndhoanit1112/go-api-clean-architecture/internal/presentation/http"
)

// Configs struct handles all dependencies required for handling configurations
type Configs struct {
}

// GetHTTPConfig returns the configuration required for HTTP package
func (cfg *Configs) GetHTTPConfig() (*http.Config, error) {
	readTimeout, err := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	if err != nil {
		readTimeout = 5
	}
	writeTimeout, err := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	if err != nil {
		writeTimeout = 5
	}

	return &http.Config{
		Port:         os.Getenv("PORT"),
		ReadTimeout:  time.Second * time.Duration(readTimeout),
		WriteTimeout: time.Second * time.Duration(writeTimeout),
	}, nil
}

// GetDbConfig returns the configuration required for db package
func (cfg *Configs) GetDbConfig() (*db.Config, error) {
	return &db.Config{
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		Username:  os.Getenv("DB_USERNAME"),
		Password:  os.Getenv("DB_PASSWORD"),
		DbName:    os.Getenv("DB_NAME"),
		DbCharset: os.Getenv("DB_CHARSET"),
	}, nil
}

// GetAllConfig returns an instance of Config with all the required dependencies initialized
func GetAllConfig() (*Configs, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Configs{}, nil
}
