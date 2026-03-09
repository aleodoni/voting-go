// Package transaction provides an abstraction for handling database transactions.
package transaction

import "context"

type Transactor interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
