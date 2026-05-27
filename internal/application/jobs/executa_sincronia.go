package jobs

import (
	"context"

	domainSincronia "github.com/aleodoni/voting-go/internal/domain/sincronia"
)

type ExecutaSincroniaJobUseCase struct {
	sincroniaRerepo domainSincronia.SincroniaRepository
}

// NewExecutaSincroniaJobUseCase cria uma nova instância de [ExecutaSincroniaJobUseCase].
func NewExecutaSincroniaJobUseCase(
	sincroniaRepo domainSincronia.SincroniaRepository,
) *ExecutaSincroniaJobUseCase {
	return &ExecutaSincroniaJobUseCase{
		sincroniaRerepo: sincroniaRepo,
	}
}

func (uc *ExecutaSincroniaJobUseCase) Execute(
	ctx context.Context,
) (*domainSincronia.Sincronia, error) {

	sincronia, err := uc.sincroniaRerepo.Sync(ctx)
	if err != nil {
		return nil, err
	}

	return sincronia, nil
}
