package config

import "os"

type Config struct {
	Port           string
	ExternalAPIURL string
}

func LoadConfig() *Config {
	return &Config{
		Port:           getEnv("PORT", "8080"),
		ExternalAPIURL: getEnv("EXTERNAL_API_URL", "https://dragonball-api.com/api"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
