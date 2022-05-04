package dataresolver

import (
	"context"
	"path/filepath"

	"github.com/opctl/opctl/sdks/go/data"
	"github.com/opctl/opctl/sdks/go/data/fs"
	"github.com/opctl/opctl/sdks/go/opspec"
)

// DataResolver resolves packages
type DataResolver interface {
	Resolve(
		ctx context.Context,
		dataRef string,
	) (data.DataHandle, error)
}

func New(dataProvider data.DataProvider, basePath string) DataResolver {
	return _dataResolver{
		dataProvider: dataProvider,
		basePath:     basePath,
	}
}

type _dataResolver struct {
	dataProvider data.DataProvider
	basePath     string
}

func (dtr _dataResolver) Resolve(
	ctx context.Context,
	dataRef string,
) (data.DataHandle, error) {
	fsProvider := fs.New(
		filepath.Join(dtr.basePath, opspec.DotOpspecDirName),
		dtr.basePath,
	)

	opDirHandle, err := data.Resolve(
		ctx,
		dataRef,
		fsProvider,
		dtr.dataProvider,
	)

	return opDirHandle, err
}
