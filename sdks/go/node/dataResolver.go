package node

import (
	"context"
	"path"

	"github.com/opctl/opctl/sdks/go/model"
)

func (core) Label() string {
	return "opctl node"
}

func (np core) TryResolve(
	ctx context.Context,
	dataRef string,
) (model.DataHandle, error) {
	if _, err := np.ListDescendants(
		ctx,
		model.ListDescendantsReq{
			DataRef: dataRef,
		},
	); err != nil {
		return nil, err
	}

	return newHandle(np, dataRef), nil
}

func newHandle(
	node Node,
	dataRef string,
) model.DataHandle {
	return handle{
		node:    node,
		dataRef: dataRef,
	}
}

// handle allows interacting w/ data sourced from an opctl node
type handle struct {
	node    Node
	dataRef string
}

func (nh handle) GetContent(
	ctx context.Context,
	contentPath string,
) (
	model.ReadSeekCloser,
	error,
) {
	return nh.node.GetData(
		ctx,
		model.GetDataReq{
			DataRef: path.Join(nh.dataRef, contentPath),
		},
	)
}

func (nh handle) ListDescendants(
	ctx context.Context,
) (
	[]*model.DirEntry,
	error,
) {
	return nh.node.ListDescendants(
		ctx,
		model.ListDescendantsReq{
			DataRef: nh.dataRef,
		},
	)
}

func (handle) Path() *string {
	return nil
}

func (nh handle) Ref() string {
	return nh.dataRef
}
