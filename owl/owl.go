// Package owl provides an implementation of a file server for the Bird
// protocol.
package owl

import (
	"errors"
	"io/fs"

	"github.com/patrickmcnamara/bird/bl"
	"github.com/patrickmcnamara/bird/seed"
)

// FileServer is a seed.Handler that can serve Seed documents from a filesystem.
//
// FS is the filesystem used and ErrHn is used to handle errors.
type FileServer struct {
	FS    fs.FS
	ErrHn ErrorHandler
}

// NewFileServer creates a new FileServer for fsys registering errHn as the
// error handler. errHn can be nil.
func NewFileServer(fsys fs.FS, errHn ErrorHandler) (fsrv *FileServer, err error) {
	if fsys == nil {
		err = errors.New("owl: NewFileServer: nil filesystem")
		return
	}
	fsrv = &FileServer{FS: fsys, ErrHn: errHn}
	return
}

// ServeBird serves the requested Bird page from a filesystem Seed document.
//
// Only files that end in seed.Extension are served. However, request Bird links
// must omit the extension. If the BL path refers to a directory, a file with
// the name of seed.Extension in that directory will be served if it exists.
//
// Filesystem:
//
//	abc/
//	abc/.sd
//	abc/1.sd
//	xyz/
//	xyz/.sd
//	xyz.sd
//
// For this filesystem, where * is the host, ServeBird would respond as such:
//
//	bird://*/abc       ->  abc/.sd
//	bird://*/abc/1     ->  abc/1.sd
//	bird://*/xyz       ->  xyz.sd
//
// If an error occurs opening a requested file, ErrHn is called using the error
// and b and sw from the request and the error that occured. If ErrHn is nil,
// the error is skipped and no response is made.
func (fsrv *FileServer) ServeBird(b bl.BL, sw *seed.Writer) {
	// open correct file
	f, err := fsrv.open(b.Path)
	// check for errors with file opening
	if err != nil {
		fsrv.errHn(b, sw, err)
		return
	}
	// serve file
	sr := seed.NewReader(f)
	seed.Copy(sw, sr)
}

// open opens a file given a path, including logic described in ServeBird. It
// purposefully does not conform to fs.FS.
func (fsrv *FileServer) open(name string) (f fs.File, err error) {
	// bad file but continue  B)
	bfbc := func(fi fs.FileInfo, err error) bool {
		return errors.Is(err, fs.ErrNotExist) || err == nil && fi.IsDir()
	}
	// possible paths, i.e. hello.sd and hello/.sd
	pth1, pth2 := paths(name)
	// find a good file if it exists
	pth := pth1
	fi, err := fs.Stat(fsrv.FS, pth)
	if bfbc(fi, err) {
		pth = pth2
		fi, err = fs.Stat(fsrv.FS, pth)
		if bfbc(fi, err) {
			err = fs.ErrNotExist
			return
		}
	}
	// open file and return
	f, err = fsrv.FS.Open(pth)
	return
}

// errHn calls ErrHn if it's not nil.
func (fsrv *FileServer) errHn(b bl.BL, sw *seed.Writer, err error) {
	if fsrv.ErrHn != nil {
		fsrv.ErrHn(b, sw, err)
	}
}
