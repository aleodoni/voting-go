package database

import (
	"fmt"
	"log"
	"time"

	"github.com/aleodoni/voting-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect estabelece uma conexão com o banco de dados PostgreSQL e configura o pool de conexões.
//
// Comportamento:
//   - utiliza as configurações fornecidas em [config.Config] para montar o DSN
//   - configura o pool com no máximo 20 conexões abertas, 5 ociosas e tempo de vida de 1 hora
func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMODE,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("erro ao acessar pool de conexões: %w", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Banco conectado com sucesso 🚀")

	return db, nil
}
