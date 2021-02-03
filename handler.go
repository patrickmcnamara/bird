package bird

import (
	"net/url"

	"github.com/patrickmcnamara/bird/seed"
)

// Handler responds to a Bird request.
type Handler interface {
	ServeBird(u *url.URL, sw *seed.Writer)
}

// HandlerFunc is a Handler that allows you to use functions to handle Bird
// requests.
type HandlerFunc func(u *url.URL, sw *seed.Writer)

// ServeBird calls the handler function.
func (hf HandlerFunc) ServeBird(u *url.URL, sw *seed.Writer) {
	hf(u, sw)
}
