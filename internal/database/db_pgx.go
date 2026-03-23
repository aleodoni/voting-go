package database

import (
	"context"
	"fmt"
	"log"

	"github.com/aleodoni/voting-go/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPGX(cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := BuildDSN(cfg)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar pgx: %w", err)
	}

	log.Println("PGX conectado com sucesso 🚀")

	return pool, nil
}
