package bird

import (
	"net/url"
)

// parseURL parses and validates a URL such that it conforms with Bird's
// requirements.
func parseURL(rawurl string) (u *url.URL, err error) {
	u, err = url.Parse(rawurl)
	if err != nil {
		return
	}
	if u.Scheme != "" {
		err = ErrProtocolRelativeURL
		return
	}
	if u.Host == "" {
		err = ErrNoHostURL
		return
	}
	return
}
