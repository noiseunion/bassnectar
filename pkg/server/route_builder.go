package server

import (
	"net/http"
)

// Route that we want to configure for our server
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// RouteBuilder provides the means to build your applications routes
type RouteBuilder struct {
	Routes []Route
}

// NewRouteBuilder creates a new one
func NewRouteBuilder() *RouteBuilder {
	return &RouteBuilder{}
}

// Register a new route with the builder
func (rb *RouteBuilder) Register(m, p string, h http.HandlerFunc) {
	r := Route{
		Method:  m,
		Path:    p,
		Handler: h,
	}

	rb.Routes = append(rb.Routes, r)
}

// RenderRoutes will attach our routes to the Server Instance
func (rb *RouteBuilder) RenderRoutes(i *Instance) {
  for _, route := range rb.Routes {
    switch route.Method {
		case http.MethodGet:
			i.router.Get(route.Path, route.Handler)
		case http.MethodPost:
			i.router.Post(route.Path, route.Handler)
		case http.MethodDelete:
			i.router.Delete(route.Path, route.Handler)
		default:
			i.Logger.Errorf("Unknown route method: %s for %+v", route.Method, route)
		} 
  }
}
