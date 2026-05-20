package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RunStagingCron(pool *pgxpool.Pool) error {
	content, err := os.ReadFile("seeds/reset_staging.sql")
	if err != nil {
		return err
	}

	_, err = pool.Exec(context.Background(), string(content))
	if err != nil {
		return err
	}

	log.Println("Staging cron configurado com sucesso ⏰")

	return nil
}
