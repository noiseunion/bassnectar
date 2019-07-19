package env

import (
	"context"

	"github.com/heetch/confita"
	"github.com/sirupsen/logrus"
)

// LoadConfig ENV vars into a specific configuration struct
func LoadConfig(cfg interface{}) {
	loader := confita.NewLoader()
	err := loader.Load(context.Background(), cfg)
	if err != nil {
		logrus.WithError(err).Info("Error loading configuration")
	}
}
