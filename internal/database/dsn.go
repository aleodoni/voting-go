package database

import (
	"fmt"

	"github.com/aleodoni/voting-go/internal/config"
)

func BuildDSN(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&TimeZone=UTC",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMODE,
	)
}
