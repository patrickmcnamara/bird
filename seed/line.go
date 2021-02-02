package seed

// Text is a text line. It is a string.
type Text string

// Header is a header line. It has a header level and text.
type Header struct {
	Level int
	Text  string
}

// Link is a link line. It has text and a URL.
type Link struct {
	Text string
	URL  string
}

// Quote is a quote line and the start or end of a quote block.
type Quote struct{}

// Code is a quote line and the start or end of a code block.
type Code struct{}

// Break is a break line. It is blank.
type Break struct{}
