// Command eagle makes Bird protocol requests.
//
// eagle requests each Bird page from each BL given as command line arguments
// and prints the Bird page. If there is an error, eagle prints the error and
// exits the program with status 1.
package main

import (
	"fmt"
	"os"

	"github.com/patrickmcnamara/bird"
	"github.com/patrickmcnamara/bird/seed"
)

func main() {
	sw := seed.NewWriter(os.Stdout)
	for _, rawBL := range os.Args[1:] {
		sr, c, err := bird.Fetch(rawBL)
		chk(err)
		err = seed.Copy(sw, sr)
		chk(err)
		err = c()
		chk(err)
	}
}

func chk(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "eagle: "+err.Error())
		os.Exit(1)
	}
}
