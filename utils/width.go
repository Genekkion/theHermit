package utils

import (
	"github.com/mattn/go-runewidth"
	"github.com/muesli/ansi"
)

const (
	defaultBufferLength = 8
)

// SplitColumns splits a string into columns based on ANSI escape sequences.
// The code is adapted from github.com/muesli/reflow/ansi/buffer.go.
func SplitColumns(s string) (columns []string) {
	buf := make([]rune, 0, defaultBufferLength)

	flag := false
	for _, r := range s {
		if r == ansi.Marker {
			flag = true
			buf = append(buf, r)
		} else if flag {
			buf = append(buf, r)
			if ansi.IsTerminator(r) {
				flag = false
			}
		} else {
			width := runewidth.RuneWidth(r)
			buf = append(buf, r)
			if width > 0 {
				columns = append(columns, string(buf))
				buf = buf[:0]
			}
		}
	}

	return columns
}
