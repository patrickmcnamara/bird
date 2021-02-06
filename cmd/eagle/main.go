package main

import (
	"fmt"
	"os"

	"github.com/patrickmcnamara/bird"
	"github.com/patrickmcnamara/bird/seed"
)

func main() {
	sw := seed.NewWriter(os.Stdout)
	for i, rawurl := range os.Args[1:] {
		sr, c, err := bird.Fetch(rawurl)
		chk(err)
		err = seed.Copy(sw, sr)
		chk(err)
		err = c()
		chk(err)
		if i != len(os.Args)-2 {
			fmt.Println()
		}
	}
}

func chk(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "owl: "+err.Error())
		os.Exit(1)
	}
}
