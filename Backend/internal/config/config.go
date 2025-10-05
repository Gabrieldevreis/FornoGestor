package config

import "os"

type Config struct {
	DBHost           string
	DBPort           string
	DBName           string
	DBUser           string
	DBPassword       string
	JWTSecret        string
	JWTRefreshSecret string
	Port             string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBName:           getEnv("DB_NAME", "fornogestor"),
		DBUser:           getEnv("DB_USER", "fornogestor"),
		DBPassword:       getEnv("DB_PASSWORD", "fornogestor123"),
		JWTSecret:        getEnv("JWT_SECRET", "your-secret-key"),
		JWTRefreshSecret: getEnv("JWT_REFRESH_SECRET", "your-refresh-secret"),
		Port:             getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
