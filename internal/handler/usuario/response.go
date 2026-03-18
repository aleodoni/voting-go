package usuario

// MeResponse representa os dados do usuário autenticado
//
//	@name	MeResponse
type MeResponse struct {
	ID           string              `json:"id"            example:"cls1abc123"`
	KeycloakID   string              `json:"keycloak_id"   example:"a1b2c3d4-..."`
	Username     string              `json:"username"      example:"joao.silva"`
	Email        string              `json:"email"         example:"joao@example.com"`
	Nome         string              `json:"nome"          example:"João Silva"`
	NomeFantasia *string             `json:"nome_fantasia" example:"Vereador João"`
	Credencial   *CredencialResponse `json:"credencial"`
}

// CredencialResponse representa os dados de uma credencial
//
//	@name	CredencialResponse
type CredencialResponse struct {
	ID              string `json:"id"               example:"cls1xyz456"`
	Ativo           bool   `json:"ativo"            example:"true"`
	PodeVotar       bool   `json:"pode_votar"       example:"false"`
	PodeAdministrar bool   `json:"pode_administrar" example:"false"`
}

// UsuarioResponse representa os dados de um usuário retornados pela API
//
//	@name	UsuarioResponse
type UsuarioResponse struct {
	ID           string              `json:"id"             example:"cls1abc123"`
	KeycloakID   string              `json:"keycloak_id"    example:"a1b2c3d4-e5f6-7890-abcd-ef1234567890"`
	Username     string              `json:"username"       example:"joao.silva"`
	Email        string              `json:"email"          example:"joao@example.com"`
	Nome         string              `json:"nome"           example:"João Silva"`
	NomeFantasia *string             `json:"nome_fantasia"  example:"Vereador João"`
	Credencial   *CredencialResponse `json:"credencial"`
}

// ListUsuariosResponse representa os dados de uma listagem de usuários retornados pela API
//
//	@name	ListUsuariosResponse
type ListUsuariosResponse struct {
	Usuarios []*UsuarioResponse `json:"usuarios"`
	Total    int64              `json:"total"  example:"42"`
	Page     int                `json:"page"   example:"1"`
	Limit    int                `json:"limit"  example:"20"`
}

// ErrorResponse representa os dados de um erro retornado pela API
//
//	@name	ErrorResponse
type ErrorResponse struct {
	Error string `json:"error" example:"mensagem de erro"`
}
