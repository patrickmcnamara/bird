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
