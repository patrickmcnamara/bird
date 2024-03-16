// Package bl implements Bird links.
package bl

const (
	MaxSize = 1024
)

// BL represents a Bird link to a Bird page.
// It is a protocol-relative URL to a Bird page.
//
// The form is:
//
//	//host/[path]
//
// host is the host of the Bird page, and path is the path. host must only
// contain a valid hostname or a IP address. paths must only contain valid path
// characters and not contain empty path tokens.
type BL struct {
	Host string // host
	Path string // path
}

// Parse parses a BL from a string.
func Parse(rawBL string) (b BL, err error) {
	return parse(rawBL)
}

// String returns the string representation of BL.
func (b BL) String() string {
	return "//" + b.Host + "/" + b.Path
}
