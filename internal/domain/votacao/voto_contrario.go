package votacao

import "time"

type VotoContrario struct {
	ID        string
	IDTexto   int
	VotoID    string
	ParecerID string
	CreatedAt time.Time
	UpdatedAt time.Time

	Parecer *Parecer
}
