package node

import (
	"context"
	"fmt"

	"github.com/opctl/opctl/sdks/go/data"
	"github.com/opctl/opctl/sdks/go/model"
)

func (c core) ListDescendants(
	ctx context.Context,
	req model.ListDescendantsReq,
) (
	[]*data.DirEntry,
	error,
) {
	if req.DataRef == "" {
		return []*data.DirEntry{}, fmt.Errorf(`"" not a valid data ref`)
	}

	dataHandle, err := c.ResolveData(ctx, req.DataRef)
	if err != nil {
		return nil, err
	}

	return dataHandle.ListDescendants(ctx)
}
