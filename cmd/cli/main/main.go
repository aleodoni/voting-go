package main

import (
	"log"
	"os"

	"github.com/aleodoni/voting-go/internal/config"
	"github.com/aleodoni/voting-go/internal/database"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: voting-cli <migrate|seed|fdw|fdw-drop>")
	}

	cfg := config.LoadConfig()

	switch os.Args[1] {
	case "migrate":
		if err := database.RunMigrations(cfg); err != nil {
			log.Fatal(err)
		}

	case "seed":
		if cfg.AppEnv != "development" && cfg.AppEnv != "staging" {
			log.Fatal("seed disabled in production")
		}
		db, err := database.ConnectPGX(cfg)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		if err := database.RunSeed(db); err != nil {
			log.Fatal(err)
		}

	case "fdw":
		if err := database.RunFDW(cfg); err != nil {
			log.Fatal(err)
		}

	case "fdw-drop":
		if err := database.RunFDWDrop(cfg); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatalf("unknown command: %s — usage: voting-cli <migrate|seed|fdw>", os.Args[1])
	}
}
