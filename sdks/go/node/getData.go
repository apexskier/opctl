package node

import (
	"context"
	"fmt"

	"github.com/opctl/opctl/sdks/go/model"
)

func (c core) GetData(
	ctx context.Context,
	eventChannel chan model.Event,
	callID string,
	req model.GetDataReq,
) (
	model.ReadSeekCloser,
	error,
) {
	if req.DataRef == "" {
		return nil, fmt.Errorf(`"" not a valid data ref`)
	}

	dataHandle, err := c.ResolveData(ctx, eventChannel, callID, req.DataRef)
	if err != nil {
		return nil, err
	}

	return dataHandle.GetContent(ctx, eventChannel, callID, "")
}
