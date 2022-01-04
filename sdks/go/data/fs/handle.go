package fs

import (
	"context"
	"os"
	"path/filepath"

	"github.com/opctl/opctl/sdks/go/data"
)

func newHandle(
	path string,
) data.DataHandle {
	return handle{
		path: path,
	}
}

// handle allows interacting w/ data sourced from the filesystem
type handle struct {
	path string
}

func (lh handle) GetContent(
	ctx context.Context,
	contentPath string,
) (
	data.ReadSeekCloser,
	error,
) {
	return os.Open(filepath.Join(lh.path, contentPath))
}

func (lh handle) ListDescendants(ctx context.Context) (
	[]*data.DirEntry,
	error,
) {
	return lh.rListDescendants(lh.path)
}

// rListDescendants recursively lists descendants of the current data node
func (lh handle) rListDescendants(
	path string,
) (
	[]*data.DirEntry,
	error,
) {
	childDirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var contents []*data.DirEntry
	for _, childDirEntry := range childDirEntries {

		absContentPath := filepath.Join(path, childDirEntry.Name())

		if childDirEntry.IsDir() {
			// recurse into child dirs
			childContents, err := lh.rListDescendants(absContentPath)
			if err != nil {
				return nil, err
			}
			contents = append(contents, childContents...)
		}

		relContentPath, err := filepath.Rel(lh.path, absContentPath)
		if err != nil {
			return nil, err
		}

		childFileInfo, err := childDirEntry.Info()
		if err != nil {
			return nil, err
		}

		contents = append(
			contents,
			&data.DirEntry{
				Mode: childFileInfo.Mode(),
				Path: filepath.Join(string(os.PathSeparator), relContentPath),
				Size: childFileInfo.Size(),
			},
		)

	}

	return contents, err
}

func (lh handle) Path() *string {
	return &lh.path
}

func (lh handle) Ref() string {
	return lh.path
}
