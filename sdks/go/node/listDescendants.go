package node

import (
	"context"
	"fmt"

	"github.com/opctl/opctl/sdks/go/model"
)

func (c core) ListDescendants(
	ctx context.Context,
	eventChannel chan model.Event,
	callID string,
	req model.ListDescendantsReq,
) (
	[]*model.DirEntry,
	error,
) {
	if req.DataRef == "" {
		return []*model.DirEntry{}, fmt.Errorf(`"" not a valid data ref`)
	}

	dataHandle, err := c.ResolveData(ctx, eventChannel, callID, req.DataRef)
	if err != nil {
		return nil, err
	}

	return dataHandle.ListDescendants(ctx, eventChannel, callID)
}
