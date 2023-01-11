package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ndhoanit1112/go-api-clean-architecture/internal/presentation/handlers"
	"gorm.io/gorm"
)

// HTTP struct holds all the dependencies required for starting HTTP server
type HTTP struct {
	server *http.Server
	cfg    *Config
}

func (s *HTTP) Start() {
	s.server.ListenAndServe()
}

// Config holds all the configuration required to start the HTTP server
type Config struct {
	Host              string
	Port              string
	TemplatesBasePath string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	DialTimeout       time.Duration
	MaxHeaderBytes    int
}

func InitServer(c *Config, db *gorm.DB) (*HTTP, error) {
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	handlers.RegisterRoutes(router, db)

	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", c.Host, c.Port),
		Handler:        router,
		ReadTimeout:    c.ReadTimeout,
		WriteTimeout:   c.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	http := &HTTP{
		server: s,
		cfg:    c,
	}

	return http, nil
}
