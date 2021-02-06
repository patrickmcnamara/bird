// Command eagle makes Bird protocol requests.
//
// eagle requests each URL given as command line arguments and prints the
// responses. If there is an error, eagle logs the error and exits the program
// with status 1.
package main

import (
	"fmt"
	"os"

	"github.com/patrickmcnamara/bird"
	"github.com/patrickmcnamara/bird/seed"
)

func main() {
	sw := seed.NewWriter(os.Stdout)
	for _, rawurl := range os.Args[1:] {
		sr, c, err := bird.Fetch(rawurl)
		chk(err)
		err = seed.Copy(sw, sr)
		chk(err)
		err = c()
		chk(err)
	}
}

func chk(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "owl: "+err.Error())
		os.Exit(1)
	}
}
