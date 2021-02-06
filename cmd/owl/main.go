// Command owl starts a basic Bird protocol file server using owl.FileServer for
// the directory it is run in.
//
// Only files with the extension seed.Extension are served, as described in the
// docs for owl.FileServer.
//
// If a requested file does not exist, owl will respond to the Bird request
// saying as much. If there is a different file I/O error, owl will respond
// saying that the file could not be served and log the error.
package main

import (
	"errors"
	"io/fs"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/patrickmcnamara/bird"
	"github.com/patrickmcnamara/bird/owl"
	"github.com/patrickmcnamara/bird/seed"
)

func main() {
	fsrv, _ := owl.NewFileServer(os.DirFS("."), errHn)
	if err := bird.Serve(":"+strconv.Itoa(int(bird.DefaultPort)), fsrv); err != nil {
		log.Fatalln(err)
	}
}

func errHn(u *url.URL, sw *seed.Writer, err error) {
	sw.Header(1, "ERROR")
	sw.Text("")
	sw.Code()
	sw.Text(u.String())
	sw.Code()
	sw.Text("")
	if errors.Is(err, fs.ErrNotExist) {
		sw.Text("File does not exist.")
	} else {
		sw.Text("File couldn't be opened.")
		log.Printf("%s: %s", u, err)
	}
}
