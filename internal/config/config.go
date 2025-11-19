package config

import (
	"os"
)

type Config struct {
	Port       string
	AuthSvcURL string
	JWTSecret  string
}

func Load() *Config {
	return &Config{
		Port:       os.Getenv("PORT"),
		AuthSvcURL: os.Getenv("AUTH_SVC_URL"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}
