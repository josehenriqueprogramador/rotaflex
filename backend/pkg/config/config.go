package config

import "os"

type Config struct {
	Port string
	DB   string
	Redis string
}

func Load() Config {
	return Config{
		Port:  getEnv("PORT", "8000"),
		DB:    getEnv("DB_URL", "postgres://user:password@db:5432/rotasflex"),
		Redis: getEnv("REDIS_URL", "redis://redis:6379"),
	}
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
