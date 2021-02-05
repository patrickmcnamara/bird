package seed

import (
	"fmt"
	"io"
	"strings"
)

// Writer writes documents in the Seed format.
//
// Doc:
//	# Lorem ipsum
//
//	>>>
//	Lorem ipsum dolor sit amet,
//	consectetur adipiscing elit,
//	sed do eiusmod tempor incididunt
//	ut labore et dolore magna aliqua.
//	>>>
//
//	=> Lorem ipsum ||| https://en.wikipedia.org/wiki/Lorem_ipsum
//
// Code:
//	sw := seed.NewWriter(...)
//	sw.Header(1, "Lorem ipsum")
//	sw.Break()
//	sw.Quote()
//	sw.Text("Lorem ipsum dolor sit amet,")
//	sw.Text("consectetur adipiscing elit,")
//	sw.Text("sed do eiusmod tempor incididunt")
//	sw.Text("ut labore et dolore magna aliqua.")
//	sw.Quote()
//	sw.Break()
//	sw.Link("Lorem ipsum", "https://en.wikipedia.org/wiki/Lorem_ipsum")
//
// As can be seen above, the Writer writes a single line every time one of it's
// line methods is called.
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
	if sw.InBlock() {
		err = ErrBadLine
		return
	}
	if level < 1 || level > 6 {
		err = ErrBadHeader
		return
	}
	_, err = fmt.Fprintf(sw.w, "%s %s\n", strings.Repeat("#", level), txt)
	return
}

// Link writes a link line to w, with a given text and URL.
func (sw *Writer) Link(txt, url string) (err error) {
	if sw.InBlock() {
		err = ErrBadLine
		return
	}
	_, err = fmt.Fprintf(sw.w, "=> %s [%s]\n", txt, url)
	return
}

// Quote writes a quote marker line to w. All following text lines are part of
// the quote. The quote marker indicates the start and end of a quote block. It
// is invalid to put anything other than text lines in a quote block and all
// other line methods will return ErrBadLine in this case.
func (sw *Writer) Quote() (err error) {
	if sw.c {
		err = ErrBadLine
		return
	}
	sw.q = !sw.q
	_, err = fmt.Fprintln(sw.w, ">>>")
	return
}

// Code writes a code marker line to w. All following text lines are part of the
// code block. The code marker indicates the start and end of a code block. It
// is invalid to put anything other than text lines in a code block and all
// other line methods will return ErrBadLine in this case.
func (sw *Writer) Code() (err error) {
	if sw.q {
		err = ErrBadLine
		return
	}
	sw.c = !sw.c
	_, err = fmt.Fprintln(sw.w, "```")
	return
}

// Line adds a line to w. The type of this line is determined by the given line.
// If line is a not a valid line type, err is ErrBadLine.
func (sw *Writer) Line(line interface{}) (err error) {
	switch l := line.(type) {
	case Text:
		err = sw.Text(string(l))
	case Header:
		err = sw.Header(l.Level, l.Text)
	case Link:
		err = sw.Link(l.Text, l.URL)
	case Quote:
		err = sw.Quote()
	case Code:
		err = sw.Code()
	default:
		err = ErrInvalidLine
	}
	return
}

// InBlock tests whether the seed document is in quote or code block.
func (sw *Writer) InBlock() bool {
	return sw.q || sw.c
}
