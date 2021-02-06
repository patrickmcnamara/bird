// Package owl provides an implementation of a file server for the Bird protocol.
package owl

import (
	"errors"
	"io/fs"
	"net/url"

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

// ServeBird serves the requested Seed document.
//
// Only files that end in seed.Extension are served. If the Bird request URL path
// does not end in seed.Extension, it will be appended before finding the file. If
// the URL path refers to a directory, a file with the name seed.Extension will
// be served if it exists.
//
// Filesystem:
//	abc
//	abc/.sd
//	abc/1.sd
//	xyz
//	xyz/.sd
//	xyz.sd
//
// For this filesystem, where * is the host, ServeBird would respond as such:
//	bird://*/abc       ->  abc/.sd
//	bird://*/abc/.sd   ->  abc/.sd
//	bird://*/abc/1     ->  abc/1.sd
//	bird://*/abc/1.sd  ->  abc/1.sd
//	bird://*/xyz       ->  xyz.sd
//
// If an error occurs opening a requested file, ErrHn is called using the error
// and u and sw from the request and the error that occured. If ErrHn is nil,
// the error is skipped and no response is made.
func (fsrv *FileServer) ServeBird(u *url.URL, sw *seed.Writer) {
	// open correct file
	f, err := fsrv.open(u.Path)
	// check for errors with file opening
	if err != nil {
		fsrv.errHn(u, sw, err)
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
func (fsrv *FileServer) errHn(u *url.URL, sw *seed.Writer, err error) {
	if fsrv.ErrHn != nil {
		fsrv.ErrHn(u, sw, err)
	}
}
