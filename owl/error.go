package owl

import (
	"github.com/patrickmcnamara/bird/bl"
	"github.com/patrickmcnamara/bird/seed"
)

// ErrorHandler is like a bird.Handler that responds to an error. It can be used
// to write a response to a Bird request and log errors, for example.
type ErrorHandler func(b bl.BL, sw *seed.Writer, err error)
