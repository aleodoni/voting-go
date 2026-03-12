// Package config provides configuration management for the application.
package config

import (
	"os"

	"github.com/aleodoni/voting-go/pkg/logger"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName        string
	AppVersion     string
	AppPort        string
	AppEnv         string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBSSLMODE      string
	JWKSURL        string
	KeycloakIssuer string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		logger.Warning("No .env file found.")
	}

	return &Config{
		AppName:        getEnv("APPNAME", "Voting API"),
		AppVersion:     getEnv("APPVERSION", "1.0.0"),
		AppPort:        getEnv("APPPORT", "8080"),
		AppEnv:         getEnv("APPENV", "development"),
		DBHost:         getEnv("DBHOST", "localhost"),
		DBPort:         getEnv("DBPORT", "15432"),
		DBUser:         getEnv("DBUSER", "postgres"),
		DBPassword:     getEnv("DBPASSWORD", "postgres"),
		DBName:         getEnv("DBNAME", "voting_db"),
		DBSSLMODE:      getEnv("DBSSLMODE", "disable"),
		JWKSURL:        getEnv("JWKSURL", "http://localhost:8081/realms/voting-realm/protocol/openid-connect/certs"),
		KeycloakIssuer: getEnv("KEYCLOAK_ISSUER", "http://localhost:8081/realms/voting-realm"),
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
