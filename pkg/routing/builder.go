package routing

import (
	"net/http"
)

// IServerInstance interface for server instances that implement route buidler
// features
type IServerInstance interface {
	RegisterRoute(r Route)
}

// IRouteBuilder interface
type IRouteBuilder interface {
	Register(m, p string, h http.HandlerFunc)
	RenderRoutes(i IServerInstance)
}

// Builder provides the means to build your applications routes
type Builder struct {
	Routes []Route
}

// NewRouteBuilder creates a new one
func NewRouteBuilder() *Builder {
	return &Builder{}
}

// Register a new route with the builder
func (b *Builder) Register(m, p string, h http.HandlerFunc) {
	r := Route{
		Method:  m,
		Path:    p,
		Handler: h,
	}

	b.append(r)
}

func (b *Builder) append(r Route) {
	b.Routes = append(b.Routes, r)
}

// RenderRoutes will attach our routes to the Server Instance
// func (rb *RouteBuilder) RenderRoutes(i *server.Instance) {
func (b Builder) RenderRoutes(i IServerInstance) {
	for _, route := range b.Routes {
		i.RegisterRoute(route)
	}
}
