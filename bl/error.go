package bl

import "errors"

var (
	// Parse errors
	ErrTooShort         = errors.New("bl: too short")
	ErrTooLong          = errors.New("bl: too long")
	ErrNoPreslashes     = errors.New("bl: missing preslashes")
	ErrNoMidslash       = errors.New("bl: missing midslash")
	ErrEmptyHost        = errors.New("bl: empty host")
	ErrIllegalCharacter = errors.New("bl: illegal character")
	ErrEmptyPathToken   = errors.New("bl: empty path token")
	ErrPostslash        = errors.New("bl: needless postslash")

	// URL errors
	ErrBadScheme = errors.New("bl: wrong URL scheme")
)
