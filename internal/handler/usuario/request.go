package usuario

type AtualizaFantasiaCredenciaisRequest struct {
	UserID      string  `json:"user_id" binding:"required"`
	DisplayName *string `json:"display_name"`
	IsActive    bool    `json:"is_active"`
	CanAdmin    bool    `json:"can_admin"`
	CanVote     bool    `json:"can_vote"`
}
