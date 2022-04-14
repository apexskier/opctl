package dataresolver

import (
	"context"
	"os"
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

func New(dataProvider data.DataProvider) DataResolver {
	return _dataResolver{
		dataProvider: dataProvider,
	}
}

type _dataResolver struct {
	dataProvider data.DataProvider
}

func (dtr _dataResolver) Resolve(
	ctx context.Context,
	dataRef string,
) (data.DataHandle, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fsProvider := fs.New(
		filepath.Join(cwd, opspec.DotOpspecDirName),
		cwd,
	)

	opDirHandle, err := data.Resolve(
		ctx,
		dataRef,
		fsProvider,
		dtr.dataProvider,
	)

	return opDirHandle, err
}
