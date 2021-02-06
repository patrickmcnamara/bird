package owl

import (
	"path"
	"strings"

	"github.com/patrickmcnamara/bird/seed"
)

func seedExt(pth string) string {
	if path.Ext(pth) == seed.Extension {
		return pth
	}
	if path.Base(pth) == "." {
		return seed.Extension
	}
	return pth + seed.Extension
}

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

func relPath(pth string) string {
	if !strings.HasSuffix(pth, "./") {
		return "./" + pth
	}
	return pth
}

func paths(pth string) (pth1, pth2 string) {
	pth = path.Clean(strings.TrimPrefix(pth, "/"))
	pth1 = seedExt(pth)
	pth2 = index(pth)
	return
}
