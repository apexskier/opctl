package node

import (
	"context"

	"github.com/opctl/opctl/sdks/go/data"
	"github.com/opctl/opctl/sdks/go/data/fs"
	"github.com/opctl/opctl/sdks/go/data/git"
)

// Resolve attempts to resolve data via local filesystem or git
// nil pullCreds will be ignored
//
// expected errs:
//  - ErrDataProviderAuthentication on authentication failure
//  - ErrDataProviderAuthorization on authorization failure
func (cr core) ResolveData(
	ctx context.Context,
	dataRef string,
) (
	data.DataHandle,
	error,
) {
	return data.Resolve(
		ctx,
		dataRef,
		fs.New(),
		git.New(cr.gitOpsDir),
	)
}
