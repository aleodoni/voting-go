package jobs

import (
	"context"

	domainJob "github.com/aleodoni/voting-go/internal/domain/job"
)

type FechaVotacoesAbertasJobUseCase struct {
	jobsRepo domainJob.JobRepository
}

// NewFechaVotacoesAbertasJobUseCase cria uma nova instância de [FechaVotacoesAbertasJobUseCase].
func NewFechaVotacoesAbertasJobUseCase(
	jobsRepo domainJob.JobRepository,
) *FechaVotacoesAbertasJobUseCase {
	return &FechaVotacoesAbertasJobUseCase{
		jobsRepo: jobsRepo,
	}
}

func (uc *FechaVotacoesAbertasJobUseCase) Execute(
	ctx context.Context,
) error {

	err := uc.jobsRepo.FecharVotacoesAbertas(ctx)
	if err != nil {
		return err
	}

	return nil
}
