// Package directry provides helpful utility functions dealing with file paths.
package directry

import "os"

// Cwd returns the full pathname of the current working directory.
// If the full pathname is unable to be resolved, then "." is returned.
func Cwd() (path string) {
	var err error
	path, err = os.Getwd()
	if err != nil {
		return "."
	}

	return
}
