package seed

import "errors"

var (
	// ErrBadLine indicates that it is illegal to insert the line. This usually
	// happens when putting a non-text line in a quote or code block.
	ErrBadLine = errors.New("seed: bad line")
	// ErrInvalidHeaderLvl indicates that the header level is invalid. A header
	// level must be between 1 and 6.
	ErrInvalidHeaderLvl = errors.New("seed: invalid header level")
	// ErrInvalidLine indicates that the given line type couldn't be converted
	// to a valid line type.
	ErrInvalidLine = errors.New("seed: invalid line type")
)
