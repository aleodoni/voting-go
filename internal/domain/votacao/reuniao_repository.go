package votacao

import (
	"context"
)

type ReuniaoRepository interface {
	FindReuniaoByID(ctx context.Context, reuniaoID string) (*Reuniao, error)
	GetReunioesDia(ctx context.Context) ([]*Reuniao, error)
}
