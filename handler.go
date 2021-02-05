package bird

import (
	"net/url"

	"github.com/patrickmcnamara/bird/seed"
)

// Handler responds to a Bird request.
//
// ServeBird writes a Seed document to sw in response to the request URL u.
type Handler interface {
	ServeBird(u *url.URL, sw *seed.Writer)
}

// HandlerFunc is a Handler that allows you to use a function to handle Bird
// requests.
//
//	f := func(u *url.URL, sw *seed.Writer) {
//		sw.Text("Hello, world!")
//	}
//	h := bird.HandlerFunc(f)
//	bird.Serve(..., h)
//
// f is type converted to a HandlerFunc here and is used as a Handler.
type HandlerFunc func(u *url.URL, sw *seed.Writer)

// ServeBird calls the handler function.
func (hf HandlerFunc) ServeBird(u *url.URL, sw *seed.Writer) {
	hf(u, sw)
}
