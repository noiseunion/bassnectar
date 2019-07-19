package env

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// LoadDotEnv will load our .env files in the desired order of precedence
// https://github.com/joho/godotenv
//
func LoadDotEnv() {
	// The .env.local file trumps all things so we load that first
	godotenv.Load(".env.local")

	// This assumes that we are in the development environment.  The
	// value of GO_ENV should be set on the host itself in order for
	// us to load other environment files.
	env := GetEnv("GO_ENV", "development")
	log.WithField("GO_ENV", env).Info("Loading environment...")
	godotenv.Load(".env." + env)

	// Load the default .env file
	godotenv.Load()
}

// GetEnv helps us by introducing a simple way to have a default fallback
// value when an ENV key is not found in the currently loaded environment
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
