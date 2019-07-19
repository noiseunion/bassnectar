package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"github.com/go-chi/cors"
)

// NewInstance creates an instance of our server
func NewInstance(cfg *Config) *Instance {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Heartbeat("/health"))
	router.Use(middleware.URLFormat)
	router.Use(middleware.Recoverer)

	logger := logrus.New()
	addr := fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port)

	if !cfg.DisableCORS {
		c := cors.New(cfg.CORS)
		router.Use(c.Handler)
	}

	return &Instance{
		Logger: logger,
		router: router,
		routeBuilder: cfg.RouteBuilder,
		
		httpServer: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
}

// Start our HTTP server up and configure it
func (instance *Instance) Start() {
	instance.RenderRoutes()

	// Now we are going to start up the HTTP Server
	// This method is blocking, so the application hangs here.
	instance.Logger.Infof("Listening on %v", instance.httpServer.Addr)

	err := instance.httpServer.ListenAndServe()

	if err != http.ErrServerClosed {
		instance.Logger.WithError(err).Error("The HTTP server has stopped unexpectedly")
		instance.Stop()
	} else {
		instance.Logger.WithError(err).Info("HTTP Server has stopped.")
	}
}

// Stop our HTTP server gracefully - if possible
func (instance *Instance) Stop() {
	instance.Logger.Info("Shutting down the HTTP Server")
}
