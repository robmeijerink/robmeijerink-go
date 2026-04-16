package config

import "os"

type AppConfig struct {
	IsProd bool
	Port   string
}

func Load() *AppConfig {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000" // Default port for dev with vite
	}

	return &AppConfig{
		IsProd: os.Getenv("APP_ENV") == "production",
		Port:   port,
	}
}
