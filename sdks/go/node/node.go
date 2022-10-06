package node

import (
	"context"

	"github.com/opctl/opctl/sdks/go/data"
	"github.com/opctl/opctl/sdks/go/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/node.go . Node

// Node is the main structure to run and interact with ops
type Node interface {
	data.DataProvider

	// StartOp starts an op and returns the root call ID
	StartOp(
		ctx context.Context,
		eventChannel chan model.Event,
		req model.StartOpReq,
	) (
		outputs map[string]*model.Value,
		err error,
	)
}
