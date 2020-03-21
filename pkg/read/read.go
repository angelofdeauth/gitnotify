// @File:     read.go
// @Created:  2020-03-21 03:20:30
// @Modified: 2020-03-21 04:52:37
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package read provides an interface for reading from data sources.
package read

import "github.com/angelofdeauth/xnotify/pkg/read/box"

// EmbeddedFile* functions are provided for reading from files embedded inside of the executable.

// EmbeddedFileToString reads the path argument and returns the contents as a string.
func EmbeddedFileToString(path string) (string, error) {
	return string(box.Get(path)), nil
}

// EmbeddedTemplateToString calls EmbeddedFileToString
func EmbeddedTemplateToString(path string) (string, error) {
	return EmbeddedFileToString(path), nil
}
