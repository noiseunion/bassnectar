package server

import "github.com/go-chi/cors"

// Config for the server
type Config struct {
	Addr         	string `config:"addr"`
	Port         	string `config:"port"`
	RouteBuilder 	*RouteBuilder
	CORS         	cors.Options
	DisableCORS  	bool
	ProtectedRoutes bool
}
