package node

import (
	"context"
	"fmt"
	"path"

	"github.com/opctl/opctl/sdks/go/data"
)

func (c core) Resolve(
	ctx context.Context,
	dataRef string,
) (data.DataHandle, error) {
	h := newHandle(c, dataRef)
	if _, err := h.ListDescendants(ctx); err != nil {
		return nil, err
	}
	return h, nil
}

func newHandle(
	node core,
	dataRef string,
) data.DataHandle {
	return handle{
		node:    node,
		dataRef: dataRef,
	}
}

// handle allows interacting w/ data sourced from an opctl node
type handle struct {
	node    core
	dataRef string
}

func (nh handle) GetContent(
	ctx context.Context,
	contentPath string,
) (
	data.ReadSeekCloser,
	error,
) {
	dataRef := path.Join(nh.dataRef, contentPath)

	if dataRef == "" {
		return nil, fmt.Errorf(`"" not a valid data ref`)
	}

	dataHandle, err := nh.node.resolveData(ctx, dataRef)
	if err != nil {
		return nil, err
	}

	return dataHandle.GetContent(ctx, "")
}

func (nh handle) ListDescendants(
	ctx context.Context,
) (
	[]*data.DirEntry,
	error,
) {
	if nh.dataRef == "" {
		return []*data.DirEntry{}, fmt.Errorf(`"" not a valid data ref`)
	}

	dataHandle, err := nh.node.resolveData(ctx, nh.dataRef)
	if err != nil {
		return nil, err
	}

	return dataHandle.ListDescendants(ctx)
}

func (handle) Path() *string {
	return nil
}

func (nh handle) Ref() string {
	return nh.dataRef
}
