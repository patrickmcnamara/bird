// Package owl provides an implementation of a file server for the Bird protocol.
package owl

import (
	"errors"
	"io/fs"
	"log"
	"net/url"

	"github.com/patrickmcnamara/bird/seed"
)

// FileServer is a seed.Handler that can serve Seed documents from a filesystem.
//
// FS is the filesystem used and Logger is the logger used for internal errors.
type FileServer struct {
	FS     fs.FS
	Logger *log.Logger
}

// NewFileServer creates a new FileServer for fsys.
func NewFileServer(fsys fs.FS) *FileServer {
	return &FileServer{FS: fsys, Logger: log.Default()}
}

// ServeBird serves the requested Seed document.
//
// If the requested filename does not end in seed.Extension, that will be
// appended to the filename before finding the file. This means that only files with
// the ".sd" extension will be served. It also means that a file named ".sd" can
// be used as an "index" for that directory. If this index file does not exist,
// it will be automatically generated.
//
// Filesystem:
//	alice
//	alice/.sd
//	bob
//	bob/hello.sd
//	bob/world.sd
//	charles
//	charles/foobar.sd
//	charles.sd
//
// In this case, FileServer would return the "alice/.sd" doc for "alice" but
// will generate an index for "bob" as it does not have a ".sd" file. "charles"
// would return the "charles.sd" doc. "bob/hello" and "bob/hello.sd" would both
// return the "bob/hello.sd" doc.
//
// If a requested file does not exist, ServeBird will return an error in response
// saying so. If a different error occurs, ServeBird will log it using Logger.
func (fsrv *FileServer) ServeBird(u *url.URL, sw *seed.Writer) {
	// try find file/dir with extension or without
	pth1, pth2 := paths(u.Path)
	pth := pth1
	f, err := fsrv.FS.Open(pth)
	if errors.Is(err, fs.ErrNotExist) {
		pth = pth2
		f, err = fsrv.FS.Open(pth)
		if errors.Is(err, fs.ErrNotExist) {
			notExist(u, sw)
			return
		} else if err != nil {
			fsrv.internalLog(u, err)
			return
		}
	} else if err != nil {
		fsrv.internalLog(u, err)
		return
	}
	fi, err := f.Stat()
	if err != nil {
		fsrv.internalLog(u, err)
		return
	}

	// handle based on whether a file or dir was found
	if fi.IsDir() {
		index(u, sw, fsrv.FS, pth)
		return
	} else if err := seed.Copy(sw, seed.NewReader(f)); err != nil {
		fsrv.internalLog(u, err)
		return
	}
}

// internalLog logs a URL and error.
func (fsrv *FileServer) internalLog(u *url.URL, err error) {
	fsrv.Logger.Printf("%s: %s", u, err)
}

// index writes the index of a directory to sw.
func index(u *url.URL, sw *seed.Writer, fsys fs.FS, pth string) (err error) {
	sw.Header(1, u.String())
	sw.Text("")
	ents, err := fs.ReadDir(fsys, pth)
	if err != nil {
		return
	}
	for _, ent := range ents {
		sw.Link(ent.Name(), relPath(ent.Name()))
	}
	if len(ents) < 1 {
		sw.Text("This directory is empty.")
	}
	return

}

// notExist writes a file does not exist error message to sw.
func notExist(u *url.URL, sw *seed.Writer) {
	sw.Header(1, "ERROR")
	sw.Text("")
	sw.Code()
	sw.Text(u.String())
	sw.Code()
	sw.Text("")
	sw.Text("This file does not exist.")
}
