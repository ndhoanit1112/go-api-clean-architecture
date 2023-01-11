package main

import (
	"github.com/ndhoanit1112/go-api-clean-architecture/internal/configs"
	"github.com/ndhoanit1112/go-api-clean-architecture/internal/infrastructure/db"
	"github.com/ndhoanit1112/go-api-clean-architecture/internal/presentation/http"
)

func main() {
	// Get all configuration
	cfg, err := configs.GetAllConfig()
	if err != nil {
		return
	}

	// Get HTTP server configuration
	httpCfg, err := cfg.GetHTTPConfig()
	if err != nil {
		return
	}

	// Get database connection configuration
	dbCfg, err := cfg.GetDbConfig()
	if err != nil {
		return
	}

	// Initialize DB instance
	dbInstance, err := db.InitDb(dbCfg)
	if err != nil {
		return
	}

	// Initialize HTTP server
	http, err := http.InitServer(httpCfg, dbInstance)
	if err != nil {
		return
	}

	// Start HTTP server
	http.Start()
}
