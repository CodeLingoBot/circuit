// Package pty provides functions for working with Unix terminals.
package pty

import (
	"os"
)

// Open: Opens a pty and its corresponding tty.
func Open() (pty, tty *os.File, err error) {
	return open()
}
