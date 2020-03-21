/*
 * @File:     read.go
 * @Created:  2020-03-19 14:53:27
 * @Modified: 2020-03-19 17:04:36
 * @Author:   Antonio Escalera
 * @Commiter: Antonio Escalera
 * @Mail:     aj@angelofdeauth.host
 * @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>
 */

/* The `read` package provides functions for reading from files.
 * EmbeddedFile* functions are provided for reading from files embedded inside of the executable.
 *
 */

package read

import (
	"io/ioutil"

	"github.com/markbates/pkger"
)

// EmbeddedFileToString reads the file argument and returns the contents as a string.
func EmbeddedFileToString(path string) (string, error) {
	f, errf := pkger.Open(path)
	if errf != nil {
		return "", errf
	}
	defer f.Close()
	b, errb := ioutil.ReadAll(f)
	if errb != nil {
		return "", errb
	}
	return string(b), nil
}
