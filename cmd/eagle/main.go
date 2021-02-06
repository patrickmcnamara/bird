package main

import (
	"fmt"
	"log"
	"os"

	"github.com/patrickmcnamara/bird"
	"github.com/patrickmcnamara/bird/seed"
)

func main() {
	sw := seed.NewWriter(os.Stdout)
	for i, rawurl := range os.Args[1:] {
		sr, c, err := bird.Fetch(rawurl)
		if err != nil {
			log.Println(err)
			continue
		}
		if err = seed.Copy(sw, sr); err != nil {
			log.Println(err)
			continue
		}
		if err := c(); err != nil {
			log.Println(err)
			continue
		}
		if i != len(os.Args)-2 {
			fmt.Println()
		}
	}
}
