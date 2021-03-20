// Code generated by "esc -pkg=opfile -o validate_schema.go -private ../../../../opspec/opfile/jsonschema.json"; DO NOT EDIT.

package opfile

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return fis[0:limit], nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/opspec/opfile/jsonschema.json": {
		name:    "jsonschema.json",
		local:   "../../../../opspec/opfile/jsonschema.json",
		size:    47771,
		modtime: 1616432652,
		compressed: `
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
`,
	},
}

var _escDirs = map[string][]os.FileInfo{}
