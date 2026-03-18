package usuario

import domainUsuario "github.com/aleodoni/voting-go/internal/domain/usuario"

func toMeResponse(u *domainUsuario.Usuario) MeResponse {
	resp := MeResponse{
		ID:           u.ID,
		KeycloakID:   u.KeycloakID,
		Username:     u.Username,
		Email:        u.Email,
		Nome:         u.Nome,
		NomeFantasia: u.NomeFantasia,
	}

	if u.Credencial != nil {
		resp.Credencial = &CredencialResponse{
			ID:              u.Credencial.ID,
			Ativo:           u.Credencial.Ativo,
			PodeVotar:       u.Credencial.PodeVotar,
			PodeAdministrar: u.Credencial.PodeAdministrar,
		}
	}

	return resp
}

func toUsuarioResponse(u *domainUsuario.Usuario) *UsuarioResponse {
	resp := &UsuarioResponse{
		ID:           u.ID,
		KeycloakID:   u.KeycloakID,
		Username:     u.Username,
		Email:        u.Email,
		Nome:         u.Nome,
		NomeFantasia: u.NomeFantasia,
	}
	if u.Credencial != nil {
		resp.Credencial = toCredencialResponse(u.Credencial)
	}
	return resp
}

func toCredencialResponse(c *domainUsuario.Credencial) *CredencialResponse {
	return &CredencialResponse{
		ID:              c.ID,
		Ativo:           c.Ativo,
		PodeVotar:       c.PodeVotar,
		PodeAdministrar: c.PodeAdministrar,
	}
}

func toListUsuariosResponse(output *domainUsuario.ListUsuario) ListUsuariosResponse {
	usuarios := make([]*UsuarioResponse, len(output.Usuarios))
	for i, u := range output.Usuarios {
		usuarios[i] = toUsuarioResponse(u)
	}
	return ListUsuariosResponse{
		Usuarios: usuarios,
		Total:    output.Total,
		Page:     output.Page,
		Limit:    output.Limit,
	}
}
