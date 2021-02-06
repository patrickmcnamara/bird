package owl

import (
	"net/url"

	"github.com/patrickmcnamara/bird/seed"
)

// ErrorHandler is like a bird.Handler that responds to an error. It can be used
// to write a response to a Bird request or log other errors for example.
type ErrorHandler func(u *url.URL, sw *seed.Writer, err error)
