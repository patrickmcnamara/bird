package seed

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// Reader reads a Seed document.
//
//	sr := seed.NewReader(...)
//	for {
//		line, err := sr.ReadLine()
//		if err == io.EOF {
//			break
//		} else if err != nil {
//			return err
//		}
//		switch l := line.(type) {
//		case seed.Header:
//			fmt.Printf("header: level=%d text=%q\n", l.Level, l.Text)
//		case seed.Text:
//			fmt.Printf("text: text=%s\n", string(l))
//		default:
//			fmt.Println("other line type")
//		}
//	}
//
// As can be seen above, the Reader reads a single line a time and details about
// the line can be gathered with a type switches or assertions.
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
	} else if !sr.q && ls == "```" { // code
		sr.c = !sr.c
		line = Code{}
	} else if sr.InBlock() {
	} else if len(ls) > 3 && ls[:3] == "=> " { // link
		parts := strings.Split(ls[3:], "|||")
		if len(parts) == 2 {
			line = Link{
				Text: strings.TrimSuffix(parts[0], " "),
				URL:  strings.TrimPrefix(parts[1], " "),
			}
		}
	} else if len(ls) > 2 && ls[0] == '#' { // header
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
		txt, ok := line.(Text)
		if !ok { // not possible ¯\_(ツ)_/¯
			err = errors.New("non text line found in block")
			return
		}
		txts = append(txts, txt)
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
