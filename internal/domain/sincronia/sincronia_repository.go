package sincronia

import "context"

type SincroniaRepository interface {
	Sync(ctx context.Context) (*Sincronia, error)
	ListLastSincronias(
		ctx context.Context,
	) ([]*Sincronia, error)
}
