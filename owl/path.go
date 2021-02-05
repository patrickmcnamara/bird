package owl

import (
	"path"
	"strings"

	"github.com/patrickmcnamara/bird/seed"
)

func clean(pth string) string {
	pth = path.Clean(pth)
	pth = strings.TrimPrefix(pth, "/")
	pth = strings.TrimSuffix(pth, "/")
	return pth
}

func withSeedExt(pth string) string {
	if pth == "." {
		pth = seed.Extension
	} else if path.Ext(pth) != seed.Extension {
		pth += seed.Extension
	}
	return pth
}

func withoutSeedExt(pth string) string {
	pth = strings.TrimSuffix(pth, seed.Extension)
	pth = clean(pth)
	if pth == "" {
		pth = "."
	}
	return pth
}

func paths(pth string) (pth1, pth2 string) {
	pth1 = clean(withSeedExt(pth))
	pth2 = clean(withoutSeedExt(pth))
	return
}

func relPath(pth string) string {
	if !strings.HasSuffix(pth, "./") {
		pth = "./" + pth
	}
	return pth
}
