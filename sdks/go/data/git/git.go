package git

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"context"
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"

	"github.com/opctl/opctl/sdks/go/data"
	"golang.org/x/sync/singleflight"
)

// singleFlightGroup is used to ensure resolves don't race across provider intances
var resolveSingleFlightGroup singleflight.Group

// New returns a data provider which sources data from git repos
func New(
	basePath string,
) data.DataProvider {
	return &_git{
		basePath: basePath,
	}
}

type _git struct {
	basePath string
}

func (gp *_git) Label() string {
	return "git"
}

func (gp *_git) Resolve(
	ctx context.Context,
	dataRef string,
) (data.DataHandle, error) {
	// attempt to resolve within singleFlight.Group to ensure concurrent resolves don't race
	handle, err, _ := resolveSingleFlightGroup.Do(
		dataRef,
		func() (interface{}, error) {
			repoPath := filepath.Join(gp.basePath, dataRef)
			handle := newHandle(repoPath, dataRef)

			// we'll mark complete clones in case we get interrupted
			completeMarkerPath := filepath.Join(gp.basePath, fmt.Sprintf(".%x", sha1.Sum([]byte(dataRef))))

			_, err := os.Stat(completeMarkerPath)
			if err == nil {
				// complete clone found
				return handle, nil
			} else if !os.IsNotExist(err) {
				// incomplete clone; blow it away
				if err := os.RemoveAll(repoPath); err != nil {
					return nil, err
				}
			}

			// attempt clone
			if err := gp.Clone(ctx, dataRef); err != nil {
				return nil, err
			}

			// mark complete
			f, err := os.Create(completeMarkerPath)
			f.Close()

			return handle, err
		},
	)
	if err != nil {
		return nil, err
	}
	return handle.(data.DataHandle), nil
}
