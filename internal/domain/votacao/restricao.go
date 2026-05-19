package votacao

import "time"

type Restricao struct {
	ID        string
	Restricao string
	VotoID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
