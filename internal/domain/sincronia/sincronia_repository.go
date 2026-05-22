package sincronia

import "context"

type SincroniaRepository interface {
	sync(ctx context.Context) (*Sincronia, error)
}
