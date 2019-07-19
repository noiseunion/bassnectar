package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// Instance is an "instance" of our server
type Instance struct {
	Logger       *logrus.Logger
	httpServer   *http.Server
	router       *chi.Mux
	routeBuilder *RouteBuilder
}

// RenderRoutes will render our RouteBuilder routes into the server instance.
func (instance *Instance) RenderRoutes() {
  instance.routeBuilder.RenderRoutes(instance)
}
