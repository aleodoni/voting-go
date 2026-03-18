package usuario

type UpdateCredencialRequest struct {
	Ativo           bool `json:"ativo"            example:"true"`
	PodeVotar       bool `json:"pode_votar"       example:"true"`
	PodeAdministrar bool `json:"pode_administrar" example:"false"`
}

type AtualizaFantasiaCredenciaisRequest struct {
	UserID      string  `json:"user_id"      example:"cls1abc123"`
	DisplayName *string `json:"display_name" example:"Vereador João"`
	IsActive    bool    `json:"is_active"    example:"true"`
	CanAdmin    bool    `json:"can_admin"    example:"false"`
	CanVote     bool    `json:"can_vote"     example:"true"`
}
