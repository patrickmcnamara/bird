package seed

import (
	"bufio"
	"io"
	"strings"
)

// Reader reads a Seed document.
type Reader struct {
	r  io.Reader
	br *bufio.Reader
	q  bool // quote block
	c  bool // code block
}

// NewReader reads a Seed document from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{r: r, br: bufio.NewReader(r)}
}

// ReadLine reads a single line from r.
func (sr *Reader) ReadLine() (line interface{}, err error) {
	ls, err := sr.br.ReadString('\n')
	ls = strings.TrimSuffix(ls, "\n")
	line = Text(ls)
	if !sr.c && ls == ">>>" { // quote
		sr.q = !sr.q
		line = Quote{}
		return
	} else if !sr.q && ls == "```" { // code
		sr.c = !sr.c
		line = Code{}
		return
	} else if sr.InBlock() {
		return
	} else if ls == "" { // break
		line = Break{}
		return
	} else if len(ls) > 3 && ls[:3] == "=> " { // link
		parts := strings.Split(ls[3:], "|||")
		if len(parts) == 2 {
			line = Link{
				Text: strings.TrimSuffix(parts[0], " "),
				URL:  strings.TrimPrefix(parts[1], " "),
			}
		}
		return
	} else if len(ls) > 2 && ls[0] == '#' {
		parts := strings.SplitN(ls, " ", 2)
		header := parts[0]
		level := 0
		for _, ch := range header {
			if ch != '#' {
				break
			}
			level++
		}
		if level > 0 && level <= 6 {
			line = Header{
				Level: level,
				Text:  parts[1],
			}
		}
		return
	}
	return
}

// ReadBlock reads the rest of the current block from r. If the Seed document is
// not in a block, it returns nil for txts and err.
func (sr *Reader) ReadBlock() (txts []Text, err error) {
	if !sr.InBlock() {
		return
	}
	for {
		var line interface{}
		line, err = sr.ReadLine()
		if err != nil && err != io.EOF {
			return
		}
		txts = append(txts, line.(Text))
		if err == io.EOF {
			break
		}
	}
	return
}

// InBlock returns whether the Seed document is currently in a block.
func (sr *Reader) InBlock() bool {
	return sr.q || sr.c
}
