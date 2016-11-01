package bundle

import (
  "github.com/opspec-io/sdk-golang/pkg/models"
)

func (this _bundle) GetOp(
opBundlePath string,
) (
opView models.OpView,
err error,
) {

  return this.opViewFactory.Construct(opBundlePath)

}
