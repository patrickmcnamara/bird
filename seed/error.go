package seed

import "errors"

var (
	// ErrBadLine indicates that it is illegal to insert the line. This usually
	// happens when putting a non-text line in a quote or code block.
	ErrBadLine = errors.New("line type cannot be added here")
	// ErrInvalidHeaderLvl indicates that the header level is invalid. A header
	// level must be between 1 and 6.
	ErrInvalidHeaderLvl = errors.New("invalid header level given")
	// ErrInvalidLine indicates that the given line type couldn't be converted
	// to a valid line type.
	ErrInvalidLine = errors.New("invalid line type given")
)
