package core

import (
	"fmt"
	"path/filepath"

	"net/url"

	"strings"

	"github.com/opctl/opctl/cli/internal/dataresolver"
	"github.com/opctl/opctl/cli/internal/nodeprovider"
	"github.com/skratchdot/open-golang/open"
)

// Uier exposes the "ui" command
type UIer interface {
	UI(
		mountRef string,
	) error
}

// newUIer returns an initialized "ui" command
func newUIer(
	dataResolver dataresolver.DataResolver,
	nodeProvider nodeprovider.NodeProvider,
) UIer {
	return _uier{
		dataResolver: dataResolver,
		nodeProvider: nodeProvider,
	}
}

type _uier struct {
	dataResolver dataresolver.DataResolver
	nodeProvider nodeprovider.NodeProvider
}

func (ivkr _uier) UI(
	mountRef string,
) error {

	if _, err := ivkr.nodeProvider.CreateNodeIfNotExists(); err != nil {
		return err
	}

	var resolvedMount string
	var err error
	if strings.HasPrefix(mountRef, ".") {
		// treat dot paths as regular rel paths
		resolvedMount, err = filepath.Abs(mountRef)
		if nil != err {
			return err
		}
	} else {
		// otherwise use same resolution as run
		mountHandle, err := ivkr.dataResolver.Resolve(
			mountRef,
			nil,
		)
		if nil != err {
			return err
		}
		resolvedMount = mountHandle.Ref()
	}

	webUIURL := fmt.Sprintf("http://localhost:42224?mount=%s", url.QueryEscape(resolvedMount))

	if err := open.Run(webUIURL); err != nil {
		return err
	}

	return nil
}
