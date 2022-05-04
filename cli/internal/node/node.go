package node

import (
	"context"
	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/node"
	"path/filepath"

	"github.com/opctl/opctl/sdks/go/data"
	"github.com/opctl/opctl/sdks/go/data/fs"
	"github.com/opctl/opctl/sdks/go/opspec"
)

func New(node node.Node, basePath string) node.Node {
	return cliNode{
		node:     node,
		basePath: basePath,
	}
}

// cliNode is a node that wraps another with some additional features for the CLI
type cliNode struct {
	node     node.Node
	basePath string
}

func (n cliNode) Label() string {
	return "cli"
}

func (n cliNode) Resolve(
	ctx context.Context,
	dataRef string,
) (data.DataHandle, error) {
	return data.Resolve(
		ctx,
		dataRef,
		// first try to resolve from the current location
		fs.New(
			filepath.Join(n.basePath, opspec.DotOpspecDirName),
			n.basePath,
		),
		// the fallback to the node being interacted with
		n.node,
	)
}

func (n cliNode) StartOp(
	ctx context.Context,
	eventChannel chan model.Event,
	req model.StartOpReq,
) (
	outputs map[string]*model.Value,
	err error,
) {
	return n.node.StartOp(ctx, eventChannel, req)
}
