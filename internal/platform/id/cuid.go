// Package id provides utilities for generating unique identifiers.
package id

import "github.com/nrednav/cuid2"

func New() string {
	return cuid2.Generate()
}
