package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/cors"
	"github.com/noiseunion/bassnectar/pkg/routing"
	"github.com/sirupsen/logrus"
)

// New creates an instance of our server
func New(cfg *Config) *Instance {
	logger := logrus.New()

	addr := fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port)

	router := routing.New()

	if !cfg.DisableCORS {
		c := cors.New(cfg.CORS)
		router.Use(c.Handler)
	}

	return &Instance{
		Logger:       logger,
		router:       router,
		routeBuilder: cfg.RouteBuilder,

		httpServer: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
}
