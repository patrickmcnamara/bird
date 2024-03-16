package bl

import "net/url"

const (
	// Scheme is the URL scheme of the Bird protocol.
	Scheme = "bird"
)

// ToURL converts a BL to a *url.URL.
func ToURL(b BL) (u *url.URL) {
	u.Scheme = Scheme
	u.Host = b.Host
	u.Path = b.Path
	return
}

// FromURL converts from a *url.URL to a BL.
func FromURL(u *url.URL) (b BL, err error) {
	if u.Scheme != Scheme {
		err = ErrBadScheme
		return
	}
	b, err = Parse("//" + u.Host + u.Path)
	return
}
