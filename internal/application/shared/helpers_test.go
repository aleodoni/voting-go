package shared_test

import (
	"github.com/aleodoni/go-ddd/domain"
	domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"
)

// adminAtivo retorna um usuário administrador ativo.
func adminAtivo(keycloakID string) *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-admin"),
		KeycloakID:    keycloakID,
		Credencial: &domainUsuario.Credencial{
			Ativo:           true,
			PodeAdministrar: true,
			PodeVotar:       true,
		},
	}
}

// usuarioSemAdmin retorna um usuário ativo mas sem permissão de administrador.
func usuarioSemAdmin(keycloakID string) *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-comum"),
		KeycloakID:    keycloakID,
		Credencial: &domainUsuario.Credencial{
			Ativo:           true,
			PodeAdministrar: false,
			PodeVotar:       true,
		},
	}
}

// usuarioSemVoto retorna um usuário ativo mas sem permissão de voto.
func usuarioSemVoto(keycloakID string) *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-sem-voto"),
		KeycloakID:    keycloakID,
		Credencial: &domainUsuario.Credencial{
			Ativo:     true,
			PodeVotar: false,
		},
	}
}

// usuarioInativo retorna um usuário com credencial inativa.
func usuarioInativo(keycloakID string) *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-inativo"),
		KeycloakID:    keycloakID,
		Credencial: &domainUsuario.Credencial{
			Ativo:           false,
			PodeAdministrar: true,
			PodeVotar:       true,
		},
	}
}

// usuarioSemCredencial retorna um usuário sem credencial cadastrada.
func usuarioSemCredencial(keycloakID string) *domainUsuario.Usuario {
	return &domainUsuario.Usuario{
		AggregateRoot: domain.NewAggregateRoot("user-sem-cred"),
		KeycloakID:    keycloakID,
		Credencial:    nil,
	}
}
