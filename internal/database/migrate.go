package database

import (
	"fmt"
	"log"

	"github.com/aleodoni/voting-go/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(cfg *config.Config) error {
	dsn := BuildDSN(cfg)

	m, err := migrate.New(
		"file://migrations",
		dsn,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration error: %w", err)
	}

	log.Println("Migrations aplicadas com sucesso 📦")

	return nil
}
