package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RunSeed(pool *pgxpool.Pool) error {
	content, err := os.ReadFile("seeds/seed.sql")
	if err != nil {
		return err
	}

	_, err = pool.Exec(context.Background(), string(content))

	log.Println("Seed executado com sucesso 🌱")

	return err
}
