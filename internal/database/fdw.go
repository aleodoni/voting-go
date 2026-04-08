package database

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aleodoni/voting-go/internal/config"
)

// RunFDW configura o Foreign Data Wrapper no PostgreSQL
func RunFDW(cfg *config.Config) error {
	files := []string{
		"fdw/01_extension.sql",
		"fdw/02_server.sql",
		"fdw/03_user_mapping.sql",
		"fdw/04_reunioes.sql",
		"fdw/05_pareceres.sql",
		"fdw/06_projetos.sql",
	}

	pool, err := ConnectPGX(cfg)
	if err != nil {
		return err
	}
	defer pool.Close()

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		sqlText := string(content)

		replacer := strings.NewReplacer(
			"${DB_SPL_HOST}", cfg.DBSPLHost,
			"${DB_SPL_NAME}", cfg.DBSPLName,
			"${DB_SPL_PORT}", cfg.DBSPLPort,
			"${DB_SPL_USER}", cfg.DBSPLUser,
			"${DB_SPL_PASSWORD}", cfg.DBSPLPassword,
		)

		sqlText = replacer.Replace(sqlText)

		_, err = pool.Exec(context.Background(), sqlText)
		if err != nil {
			return fmt.Errorf("erro em %s: %w", file, err)
		}

		fmt.Printf("✅ %s executado\n", file)
	}

	fmt.Println("🚀 FDW configurado com sucesso")
	return nil
}

func RunFDWDrop(localCfg *config.Config) error {
	pool, err := ConnectPGX(localCfg)
	if err != nil {
		return err
	}
	defer pool.Close()

	content, err := os.ReadFile("fdw/drop.sql")
	if err != nil {
		return err
	}

	_, err = pool.Exec(context.Background(), string(content))
	if err != nil {
		return err
	}

	fmt.Println("🗑️ FDW removido com sucesso")
	return nil
}
