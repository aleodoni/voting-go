package config

import (
	"os"
	"strings"

	"github.com/aleodoni/voting-go/pkg/logger"
	"github.com/joho/godotenv"
)

// Config contém todas as configurações necessárias para a execução da aplicação.
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
	AllowOrigins   []string
}

// LoadConfig carrega as configurações da aplicação a partir de variáveis de ambiente.
//
// Comportamento:
//   - tenta carregar um arquivo .env na raiz do projeto
//   - caso o arquivo não seja encontrado, emite um aviso e utiliza os valores padrão
//   - os valores padrão são adequados para execução em ambiente de desenvolvimento local
func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		logger.Warning("No .env file found.")
	}

	allowOrigins := parseOrigins(
		getEnv("ALLOW_ORIGINS", "http://localhost:5173,http://localhost:5174"),
	)

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
		AllowOrigins:   allowOrigins,
	}
}

// getEnv retorna o valor da variável de ambiente identificada por key.
// Caso a variável não esteja definida, retorna o valor de fallback.
func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func parseOrigins(origins string) []string {
	list := strings.Split(origins, ",")
	for i := range list {
		list[i] = strings.TrimSpace(list[i])
	}
	return list
}
