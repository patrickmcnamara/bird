package bird

import (
	"github.com/patrickmcnamara/bird/seed"
)

// Handler responds to a Bird request.
type Handler interface {
	ServeBird(rawurl string, sw *seed.Writer)
}

// HandlerFunc is a Handler that allows you to use functions to handle Bird
// requests.
type HandlerFunc func(rawurl string, sw *seed.Writer)

// ServeBird calls the handler function.
func (hf HandlerFunc) ServeBird(rawurl string, sw *seed.Writer) {
	hf(rawurl, sw)
}
