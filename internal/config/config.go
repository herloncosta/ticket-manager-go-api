package config

import "os"

type Config struct {
	DBPath string
}

func Load() *Config {
	return &Config{
		DBPath: getEnv("DB_PATH", "data.db"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
