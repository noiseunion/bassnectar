package main

import (
	"net/http"

	"github.com/go-chi/cors"
	"github.com/noiseunion/bassnectar/pkg/env"
	"github.com/noiseunion/bassnectar/pkg/routing"
	"github.com/noiseunion/bassnectar/pkg/server"
)

var instance *server.Instance
var config server.Config

// This file is pretty much just an example of how to use the packages
// in this library.
func init() {
	env.LoadDotEnv()

	// Load our server config
	config = server.Config{
		RouteBuilder: registerRoutes(),
		CORS: cors.Options{
			AllowedOrigins: []string{"*", "localhost"},
		},
	}

	config.LoadEnv()
}

func registerRoutes() *routing.Builder {
	b := routing.NewRouteBuilder()

	b.Register(http.MethodGet, "/foo", handleFoo)

	return b
}

func handleFoo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FOO BAR"))
}

func main() {
	i := server.New(&config)
	i.Start()
}
