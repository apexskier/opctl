package data

import (
	"context"
)

// DataProvider is the interface for something that provides data
type DataProvider interface {
	Label() string

	// Resolve resolves a package from the source.
	Resolve(
		ctx context.Context,
		dataRef string,
	) (DataHandle, error)
}
