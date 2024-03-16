package owl

import (
	"path"

	"github.com/patrickmcnamara/bird/seed"
)

// paths returns both possible index paths given pth.
func paths(pth string) (pth1, pth2 string) {
	pth1 = path.Join(pth, seed.Extension)
	pth2 = pth + seed.Extension
	return
}
