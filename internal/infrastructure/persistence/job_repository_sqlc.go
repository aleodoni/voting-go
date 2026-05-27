package persistence

import (
	"context"
	"fmt"

	"github.com/aleodoni/voting-go/internal/domain/job"
	db "github.com/aleodoni/voting-go/internal/infrastructure/persistence/sqlc/generated"
	"github.com/jackc/pgx/v5/pgxpool"
)

type jobRepositorySQLC struct {
	q *db.Queries
}

func NewJobRepositorySQLC(pool *pgxpool.Pool) job.JobRepository {
	return &jobRepositorySQLC{
		q: db.New(pool),
	}
}

func (r *jobRepositorySQLC) FecharVotacoesAbertas(ctx context.Context) error {
	err := r.q.FecharVotacoesAbertas(ctx)
	if err != nil {
		return fmt.Errorf("FecharVotacoesAbertas: %w", err)
	}
	return nil
}
