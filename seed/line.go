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

// Quote is a quote marker line.
type Quote struct{}

// Code is a code marker line.
type Code struct{}
