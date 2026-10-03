package sincronia

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	domainSincronia "github.com/aleodoni/voting-go/internal/domain/sincronia"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
)

// ExecutaSincroniaInput contém os dados necessários para executar a sincronização.
type ExecutaSincroniaInput struct {
	LoggedInUserKeycloakID string
}

// ExecutaSincroniaUseCase executa a sincronização de dados.
//
// Regras de negócio:
//   - somente usuário autenticado com permissão de administrador pode executar a sincronização
type ExecutaSincroniaUseCase struct {
	sincroniaRerepo domainSincronia.SincroniaRepository
	usuarioRepo     domainUsuario.UsuarioRepository
}

// NewExecutaSincroniaUseCase cria uma nova instância de [ExecutaSincroniaUseCase].
func NewExecutaSincroniaUseCase(
	sincroniaRepo domainSincronia.SincroniaRepository,
	usuarioRepo domainUsuario.UsuarioRepository,
) *ExecutaSincroniaUseCase {
	return &ExecutaSincroniaUseCase{
		sincroniaRerepo: sincroniaRepo,
		usuarioRepo:     usuarioRepo,
	}
}

// Autorizar verifica, sem executar a sincronização, se o usuário logado pode
// executá-la. Permite que o chamador recuse a requisição antes de iniciar uma
// execução assíncrona, cujo erro de permissão só apareceria no log.
func (uc *ExecutaSincroniaUseCase) Autorizar(ctx context.Context, input ExecutaSincroniaInput) error {
	return shared.VerificarAdmin(ctx, uc.usuarioRepo, input.LoggedInUserKeycloakID)
}

func (uc *ExecutaSincroniaUseCase) Execute(
	ctx context.Context,
	input ExecutaSincroniaInput,
) (*domainSincronia.Sincronia, error) {

	if err := uc.Autorizar(ctx, input); err != nil {
		return nil, err
	}

	sincronia, err := uc.sincroniaRerepo.Sync(ctx)
	if err != nil {
		return nil, err
	}

	return sincronia, nil
}
