package job

import "context"

type JobRepository interface {
	FecharVotacoesAbertas(ctx context.Context) error
}
