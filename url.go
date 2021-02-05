package bird

import (
	"errors"
	"net/url"
	"strings"
)

var (
	// ErrNoHostURL indicates that the URL given was missing the host.
	ErrNoHostURL = errors.New("URL should have host")
	// ErrProtocolRelativeURL indicates that the URL given was not relative to
	// the bird:// scheme.
	ErrProtocolRelativeURL = errors.New("URL should be protocol relative and have no scheme")
	// ErrTrailingSlashURL indicates that the URL given had a trailing slash,
	// which is not allowed in the Bird protocol.
	ErrTrailingSlashURL = errors.New("URL should not have trailing slash")
)

// parseURL parses and validates a URL such that it conforms with Bird's
// requirements.
func parseURL(rawurl string) (u *url.URL, err error) {
	u, err = url.Parse(rawurl)
	if err != nil {
		return
	}
	if u.Host == "" {
		err = ErrNoHostURL
		return
	}
	if u.Scheme != "" {
		err = ErrProtocolRelativeURL
		return
	}
	if strings.HasSuffix(u.Path, "/") {
		err = ErrTrailingSlashURL
		return
	}
	return
}
