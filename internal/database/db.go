// Package database provides functions for connecting to and interacting with the database.
package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"

	"github.com/aleodoni/voting-go/internal/config"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {

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
		log.Fatal("Erro ao conectar ao banco: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Erro ao acessar pool de conexões: ", err)
	}

	// Pool de conexões
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	log.Println("Banco conectado com sucesso 🚀")
}
