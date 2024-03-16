// Command owl starts a basic Bird protocol file server using owl.FileServer for
// the directory it is run in.
//
// Only files with the extension seed.Extension are served, as described in the
// docs for owl.FileServer.
//
// If a requested file does not exist or there is a permission error, owl will
// respond to the Bird request saying as much. If there is a different file I/O
// error, owl will respond saying that the file could not be served.
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"

	"github.com/patrickmcnamara/bird"
	"github.com/patrickmcnamara/bird/bl"
	"github.com/patrickmcnamara/bird/owl"
	"github.com/patrickmcnamara/bird/seed"
)

func main() {
	fsrv, _ := owl.NewFileServer(os.DirFS("."), errHn)
	err := bird.Serve(":"+strconv.Itoa(int(bird.DefaultPort)), fsrv)
	chk(err)
}

func errHn(b bl.BL, sw *seed.Writer, err error) {
	sw.Header(1, "ERROR")
	sw.Text("")
	sw.Code()
	sw.Text(b.Path)
	sw.Code()
	sw.Text("")
	if errors.Is(err, fs.ErrNotExist) || strings.Contains(err.Error(), "not a directory") {
		sw.Text("File does not exist.")
	} else if errors.Is(err, fs.ErrPermission) {
		sw.Text("File is verboten.")
	} else {
		sw.Text("File couldn't be opened.")
	}
}

func chk(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "owl: "+err.Error())
		os.Exit(1)
	}
}
