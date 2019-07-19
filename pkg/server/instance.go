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
	rb := instance.routeBuilder

	for _, route := range rb.Routes {
		switch route.Method {
		case http.MethodGet:
			instance.router.Get(route.Path, route.Handler)
		case http.MethodPost:
			instance.router.Post(route.Path, route.Handler)
		case http.MethodDelete:
			instance.router.Delete(route.Path, route.Handler)
		default:
			logrus.Errorf("Unknown route method: %s for %+v", route.Method, route)
		}
	}
}
