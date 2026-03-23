package bootstrap

import (
	"github.com/jackc/pgx/v5/pgxpool"

	persistence "github.com/aleodoni/voting-go/internal/infrastructure/persistence"
)

func buildRepositories(pgxPool *pgxpool.Pool) *repositories {
	return &repositories{
		usuario:    persistence.NewUsuarioRepositorySQLC(pgxPool),
		transactor: persistence.NewUnitOfWorkSQLC(pgxPool),
		reuniao:    persistence.NewReuniaoRepositorySQLC(pgxPool),
		votacao:    persistence.NewVotacaoRepositorySQLC(pgxPool),
	}
}
