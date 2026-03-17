// Package persistence provides helpers for working with PostgreSQL errors.
package persistence

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

// Códigos de erro do PostgreSQL
// https://www.postgresql.org/docs/current/errcodes-appendix.html
const (
	pgErrUniqueViolation     = "23505"
	pgErrForeignKeyViolation = "23503"
	pgErrNotNullViolation    = "23502"
	pgErrCheckViolation      = "23514"
)

func isPgError(err error, code string) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == code
}

func IsUniqueViolation(err error) bool {
	return isPgError(err, pgErrUniqueViolation)
}

func IsForeignKeyViolation(err error) bool {
	return isPgError(err, pgErrForeignKeyViolation)
}

func IsNotNullViolation(err error) bool {
	return isPgError(err, pgErrNotNullViolation)
}

func IsCheckViolation(err error) bool {
	return isPgError(err, pgErrCheckViolation)
}
