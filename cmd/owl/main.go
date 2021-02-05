package main

import (
	"log"
	"os"
	"strconv"

	"github.com/patrickmcnamara/bird"
	"github.com/patrickmcnamara/bird/owl"
)

func main() {
	fsrv := owl.NewFileServer(os.DirFS("."))
	if err := bird.Serve(":"+strconv.Itoa(int(bird.DefaultPort)), fsrv); err != nil {
		log.Fatalln(err)
	}
}
