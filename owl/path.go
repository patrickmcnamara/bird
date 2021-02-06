package owl

import (
	"path"
	"strings"

	"github.com/patrickmcnamara/bird/seed"
)

// seedExt adds seed.Extension to a path if it's not already there.
func seedExt(pth string) string {
	if path.Ext(pth) == seed.Extension {
		return pth
	}
	if path.Base(pth) == "." {
		return seed.Extension
	}
	return pth + seed.Extension
}

// index returns valid "index" path for the given path.
func index(pth string) string {
	if path.Base(pth) == seed.Extension {
		return pth
	}
	if path.Base(pth) == "." {
		return seed.Extension
	}
	if path.Ext(pth) == seed.Extension {
		pth = strings.TrimSuffix(pth, seed.Extension)
	}
	return path.Join(pth, seed.Extension)
}

// paths returns both valid paths given path.
func paths(pth string) (pth1, pth2 string) {
	pth = path.Clean(strings.TrimPrefix(pth, "/"))
	pth1 = seedExt(pth)
	pth2 = index(pth)
	return
}
