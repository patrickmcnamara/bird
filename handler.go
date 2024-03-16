package bird

import (
	"github.com/patrickmcnamara/bird/bl"
	"github.com/patrickmcnamara/bird/seed"
)

// Handler responds to a Bird request.
//
// ServeBird writes a Seed document to sw in response to request b.
type Handler interface {
	ServeBird(b bl.BL, sw *seed.Writer)
}

// HandlerFunc is a Handler that allows you to use a function to handle Bird
// requests.
//
//	f := func(b bl.BL, sw *seed.Writer) {
//		sw.Text("Hello, world!")
//	}
//	h := bird.HandlerFunc(f)
//	bird.Serve(..., h)
//
// f is type-converted to a HandlerFunc here and is used as a Handler.
type HandlerFunc func(b bl.BL, sw *seed.Writer)

// ServeBird calls the handler function.
func (hf HandlerFunc) ServeBird(b bl.BL, sw *seed.Writer) {
	hf(b, sw)
}
