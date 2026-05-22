package sincronia

import (
	"context"

	"github.com/aleodoni/voting-go/internal/application/shared"
	"github.com/aleodoni/voting-go/internal/domain/sincronia"
	domainSincronia "github.com/aleodoni/voting-go/internal/domain/sincronia"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
)

// RetornaSincroniasInput contém os dados necessários para executar a sincronização.
type RetornaSincroniasInput struct {
	LoggedInUserKeycloakID string
}

// RetornaSincroniasUseCase retorna as três últimas sincronizações realizadas.
//
// Regras de negócio:
//   - somente usuário autenticado com permissão de administrador pode executar
type RetornaSincroniasUseCase struct {
	sincroniaRerepo domainSincronia.SincroniaRepository
	usuarioRepo     domainUsuario.UsuarioRepository
}

// NewRetornaSincroniasUseCase cria uma nova instância de [RetornaSincroniasUseCase].
func NewRetornaSincroniasUseCase(
	sincroniaRepo domainSincronia.SincroniaRepository,
	usuarioRepo domainUsuario.UsuarioRepository,
) *RetornaSincroniasUseCase {
	return &RetornaSincroniasUseCase{
		sincroniaRerepo: sincroniaRepo,
		usuarioRepo:     usuarioRepo,
	}
}

func (uc *RetornaSincroniasUseCase) Execute(
	ctx context.Context,
	input RetornaSincroniasInput,
) (*domainSincronia.ListSincronia, error) {

	if err := shared.VerificarAdmin(ctx, uc.usuarioRepo, input.LoggedInUserKeycloakID); err != nil {
		return nil, err
	}

	sincronias, err := uc.sincroniaRerepo.ListLastSincronias(ctx)
	if err != nil {
		return nil, err
	}

	return &sincronia.ListSincronia{
		Sincronias: sincronias,
	}, nil
}
