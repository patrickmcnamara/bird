// Package seed reads and writes Seed documents.
package seed

import "io"

const (
	// Extension is the standard filename extension for Seed documents.
	Extension = ".sd"
)

// Copy copies a Seed document from sr to sw until either io.EOF is reached or
// an error occurs. err == nil on a successful Copy.
func Copy(sw *Writer, sr *Reader) (err error) {
	for {
		var line Line
		if line, err = sr.ReadLine(); err == io.EOF {
			err = nil
			return
		} else if err != nil {
			return
		}
		if err = sw.Line(line); err != nil {
			return
		}
	}
}
