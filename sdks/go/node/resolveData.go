package node

import (
	"context"
	"github.com/opctl/opctl/sdks/go/data"
	"github.com/opctl/opctl/sdks/go/data/fs"
	"github.com/opctl/opctl/sdks/go/data/git"
)

// resolveData attempts to resolve data via local filesystem or git
func (c core) resolveData(
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
		git.New(c.gitOpsDir),
	)
}
