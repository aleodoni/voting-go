// Package votacao implements the repository for managing voting sessions and votes.
package votacao

import (
	"context"
	"time"
)

type VotacaoRepository interface {
	FindReuniaoByID(ctx context.Context, reuniaoID string) (*Reuniao, error)
	GetReunioesByData(ctx context.Context, data time.Time) ([]*Reuniao, error)
}
