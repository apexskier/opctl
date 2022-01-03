package git

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/opctl/opctl/sdks/go/internal/readchunks"
	"github.com/opctl/opctl/sdks/go/model"
)

// Clone 'dataRef' to 'path'
// nil pullCreds will be ignored
//
// expected errs:
//  - ErrDataProviderAuthentication on authentication failure
//  - ErrDataProviderAuthorization on authorization failure
func (gp *_git) Clone(
	ctx context.Context,
	eventChannel chan model.Event,
	callID string,
	dataRef string,
) error {
	parsedPkgRef, err := parseRef(dataRef)
	if err != nil {
		return fmt.Errorf("invalid git ref: %w", err)
	}

	opPath := parsedPkgRef.ToPath(gp.basePath)

	url := fmt.Sprintf("https://%s", parsedPkgRef.Name)
	creds, err := getCredentials(ctx, url)
	if err != nil {
		return fmt.Errorf("invalid git ref: %w", err)
	}
	reader, writer := io.Pipe()
	defer writer.Close()
	cloneOptions := &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/tags/%s", parsedPkgRef.Version)),
		Depth:         1,
		Progress:      writer,
		Auth: &http.BasicAuth{
			Username: creds.Username,
			Password: creds.Password,
		},
	}

	// outputErr := make(chan error, 1)
	go func() {
		_ = readchunks.ReadChunks(
			reader,
			func(chunk []byte) {
				eventChannel <- model.Event{
					Timestamp: time.Now().UTC(),
					OpPullProgress: &model.OpPullProgress{
						ContainerID: callID,
						OpRef:       parsedPkgRef.String(),
						Data:        chunk,
					},
				}
			},
		)
	}()

	if _, err := git.PlainCloneContext(
		ctx,
		opPath,
		false,
		cloneOptions,
	); err != nil {
		if _, ok := err.(git.NoMatchingRefSpecError); ok {
			return fmt.Errorf("version '%s' not found", parsedPkgRef.Version)
		}
		if errors.Is(err, transport.ErrAuthenticationRequired) {
			return model.ErrDataProviderAuthentication{}
		}
		if errors.Is(err, transport.ErrAuthorizationFailed) {
			return model.ErrDataProviderAuthorization{}
		}
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) || err == ctx.Err() {
			fmt.Fprintf(os.Stderr, "cleaning up %v\n", dataRef)
			err := os.RemoveAll(opPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to cleanup partially downloaded op: %v\n", err)
			}
		}
		return err
	}

	return os.RemoveAll(filepath.Join(opPath, ".git"))
}
