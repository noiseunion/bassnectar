package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/noiseunion/bassnectar/pkg/routing"
	"github.com/sirupsen/logrus"
)

// Instance is an "instance" of our server
type Instance struct {
	Logger       *logrus.Logger
	httpServer   *http.Server
	router       *chi.Mux
	routeBuilder routing.IRouteBuilder
}

// Start our HTTP server up and configure it
func (i *Instance) Start() {
	i.routeBuilder.RenderRoutes(i)

	// Now we are going to start up the HTTP Server
	// This method is blocking, so the application hangs here.
	i.Logger.Infof("Listening on %v", i.httpServer.Addr)

	err := i.httpServer.ListenAndServe()

	if err != http.ErrServerClosed {
		i.Logger.WithError(err).Error("The HTTP server has stopped unexpectedly")
		i.Stop()
	} else {
		i.Logger.WithError(err).Info("HTTP Server has stopped.")
	}
}

// Stop our HTTP server gracefully - if possible
func (i *Instance) Stop() {
	i.Logger.Info("Shutting down the HTTP Server")
}

// RegisterRoute will attach a route to our instance
func (i *Instance) RegisterRoute(r routing.Route) {
	switch r.Method {
	case http.MethodGet:
		i.router.Get(r.Path, r.Handler)
	case http.MethodPost:
		i.router.Post(r.Path, r.Handler)
	case http.MethodDelete:
		i.router.Delete(r.Path, r.Handler)
	default:
		i.Logger.Errorf("Unknown route method: %s for %+v", r.Method, r)
	}
}
