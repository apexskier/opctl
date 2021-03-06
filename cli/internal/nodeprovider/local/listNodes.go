package local

import (
	"github.com/opctl/opctl/cli/internal/model"
	"path/filepath"
)

func (np nodeProvider) ListNodes() ([]model.NodeHandle, error) {
	pIDOfLockOwner := np.lockfile.PIdOfOwner(
		filepath.Join(
			np.dataDir.Path(),
			"pid.lock",
		),
	)
	if 0 != pIDOfLockOwner {
		nodeHandle, err := newNodeHandle(np.listenAddress)
		if nil != err {
			return nil, err
		}

		return []model.NodeHandle{
			nodeHandle,
		}, nil
	}

	return []model.NodeHandle{}, nil
}
