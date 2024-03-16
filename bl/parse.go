package bl

import "strings"

const (
	validHost = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-.:[]"
	validPath = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-./"
)

func parse(rawBL string) (b BL, err error) {
	// length
	if len(rawBL) < 2 {
		err = ErrTooShort
		return
	}
	if len(rawBL) > MaxSize {
		err = ErrTooLong
		return
	}
	// preslashes
	if rawBL[:2] != "//" {
		err = ErrNoPreslashes
		return
	}
	// host + midslash + path
	host, path, midslash := strings.Cut(rawBL[2:], "/")
	// host
	if host == "" {
		err = ErrEmptyHost
		return
	}
	for _, hch := range host {
		if !strings.ContainsRune(validHost, hch) {
			err = ErrIllegalCharacter
			return
		}
	}
	// midslash
	if !midslash {
		err = ErrNoMidslash
		return
	}
	// path
	for _, pch := range path {
		if !strings.ContainsRune(validPath, pch) {
			err = ErrIllegalCharacter
			return
		}
	}
	if strings.HasSuffix(path, "/") {
		err = ErrPostslash
		return
	}
	ptks := strings.Split(path, "/")
	if len(ptks) != 1 {
		for _, ptk := range ptks {
			if ptk == "" {
				err = ErrEmptyPathToken
				return
			}
		}
	}
	// BL
	b.Host = host
	b.Path = path
	return
}
