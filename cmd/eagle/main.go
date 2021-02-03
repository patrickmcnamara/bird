package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/patrickmcnamara/bird"
	"github.com/patrickmcnamara/bird/seed"
)

func main() {
	for _, rawurl := range os.Args[1:] {
		quote := false
		code := false
		sr, c, err := bird.Fetch(rawurl)
		if err != nil {
			log.Println(err)
			continue
		}
		for {
			line, err := sr.ReadLine()
			if err != nil {
				break
			}
			switch l := line.(type) {
			case seed.Text:
				switch {
				case quote:
					fmt.Println(">", string(l))
				case code:
					fmt.Println("`", string(l))
				default:
					fmt.Println(string(l))
				}
			case seed.Header:
				fmt.Println(strings.Repeat("#", l.Level), l.Text)
			case seed.Link:
				fmt.Printf("=> %s (%s)\n", l.Text, l.URL)
			case seed.Quote:
				quote = !quote
			case seed.Code:
				code = !code
			case seed.Break:
				fmt.Println()
			}
		}
		fmt.Println()
		if err := c(); err != nil {
			log.Println(err)
			continue
		}
	}
}
