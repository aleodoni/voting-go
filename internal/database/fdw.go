package database

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aleodoni/voting-go/internal/config"
)

func RunFDW(cfg *config.Config) error {
	if cfg.DBSPLHost == "" {
		return fmt.Errorf("configuração SPL ausente")
	}

	pool, err := ConnectPGX(cfg)
	if err != nil {
		return err
	}
	defer pool.Close()

	var serverExists bool
	var tableExists bool

	serverSQL := `
	SELECT EXISTS (
		SELECT 1
		FROM pg_foreign_server
		WHERE srvname = 'server_spl'
	)
	`

	err = pool.QueryRow(context.Background(), serverSQL).Scan(&serverExists)
	if err != nil {
		return fmt.Errorf("erro ao verificar foreign server: %w", err)
	}

	tableSQL := `
	SELECT EXISTS (
		SELECT 1
		FROM information_schema.foreign_tables
		WHERE foreign_table_schema = 'public'
		  AND foreign_table_name = 'spl_votacao_reunioes_foreign'
	)
	`

	err = pool.QueryRow(context.Background(), tableSQL).Scan(&tableExists)
	if err != nil {
		return fmt.Errorf("erro ao verificar foreign table: %w", err)
	}

	if serverExists && tableExists {
		fmt.Println("FDW já configurado completamente ✅")
		return nil
	}

	content, err := os.ReadFile("db/fdw/spl_setup.sql")
	if err != nil {
		return fmt.Errorf("erro ao ler sql fdw: %w", err)
	}

	sqlText := string(content)

	replacer := strings.NewReplacer(
		"{{DB_SPL_HOST}}", cfg.DBSPLHost,
		"{{DB_SPL_NAME}}", cfg.DBSPLName,
		"{{DB_SPL_PORT}}", cfg.DBSPLPort,
		"{{DB_SPL_USER}}", cfg.DBSPLUser,
		"{{DB_SPL_PASSWORD}}", cfg.DBSPLPassword,
	)

	sqlText = replacer.Replace(sqlText)

	_, err = pool.Exec(context.Background(), sqlText)
	if err != nil {
		return fmt.Errorf("erro ao executar fdw: %w", err)
	}

	fmt.Println("FDW configurado com sucesso 🚀")

	return nil
}
