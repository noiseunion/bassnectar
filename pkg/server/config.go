package server

import (
	"context"

	"github.com/go-chi/cors"
	"github.com/heetch/confita"
	"github.com/noiseunion/bassnectar/pkg/routing"
	"github.com/sirupsen/logrus"
)

// Config for the server
type Config struct {
	Addr         string `config:"addr"`
	Port         string `config:"port"`
	RouteBuilder routing.IRouteBuilder
	DisableCORS  bool
	CORS         cors.Options
}

// LoadEnv will load the server config from the ENV
func (c *Config) LoadEnv() {
	loader := confita.NewLoader()
	err := loader.Load(context.Background(), c)
	if err != nil {
		logrus.WithError(err).Error("Could not load the server config from the environment")
	}
}
