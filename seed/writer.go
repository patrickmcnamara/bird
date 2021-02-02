package seed

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

var (
	// ErrBadLine indicates that it is illegal to insert the line. This usually
	// happens when putting a non-text line in a quote or code block.
	ErrBadLine = errors.New("line type cannot be added here")
)

// Writer writes documents in the seed format.
type Writer struct {
	w io.Writer
	q bool // quote block
	c bool // code block
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

// Text writes a text line to w.
func (sw *Writer) Text(txt string) (err error) {
	_, err = fmt.Fprintln(sw.w, txt)
	return
}

// Header writes a header line to w, with a given header level.
func (sw *Writer) Header(level int, txt string) (err error) {
	_, err = fmt.Fprintf(sw.w, "%s %s\n", strings.Repeat("#", level), txt)
	return
}

// Link writes a link line to w, with a given text and URL.
func (sw *Writer) Link(txt, url string) (err error) {
	_, err = fmt.Fprintf(sw.w, "=> %s [%s]\n", txt, url)
	return
}

// Quote writes a quote marker line to w. All following text lines are part of
// the quote. The quote marker indicates the start and end of a quote block. It
// is invalid to put anything other than text lines in a quote block.
func (sw *Writer) Quote() (err error) {
	sw.q = !sw.q
	_, err = fmt.Fprintln(sw.w, ">>>")
	return
}

// Code writes a code marker line to w. All following text lines are part of the
// code block. The code marker indicates the start and end of a code block. It
// is invalid to put anything other than text lines in a code block.
func (sw *Writer) Code() (err error) {
	sw.c = !sw.c
	_, err = fmt.Fprintln(sw.w, "```")
	return
}

// Break adds a break line to w. This is used to space out other lines and
// separate paragraphs.
func (sw *Writer) Break() (err error) {
	_, err = fmt.Fprintln(sw.w)
	return
}

// Line adds a line to w. The type of this line is determined by the given line.
// If line is a not a valid line type, err is ErrBadLine.
func (sw *Writer) Line(line interface{}) (err error) {
	switch l := line.(type) {
	case Text:
		sw.Text(string(l))
	case Header:
		sw.Header(l.Level, l.Text)
	case Link:
		sw.Link(l.Text, l.URL)
	case Quote:
		sw.Quote()
	case Code:
		sw.Code()
	case Break:
		sw.Break()
	default:
		err = ErrBadLine
	}
	return
}

// InBlock tests whether the seed document is in quote or code block.
func (sw *Writer) InBlock() bool {
	return sw.q || sw.c
}
