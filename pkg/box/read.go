// @File:     read.go
// @Created:  2020-03-21 03:20:30
// @Modified: 2020-03-23 21:16:38
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package box

// functions provided for reading from files embedded inside of the executable.

// ReadEmbeddedFileToString reads the path argument and returns the contents as a string.
func (e *EmbedBox) ReadEmbeddedFileToString(path string) (string, error) {
	b, err := e.Get(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ReadEmbeddedTemplateToString calls EmbeddedFileToString
func (e *EmbedBox) ReadEmbeddedTemplateToString(path string) (string, error) {
	return e.ReadEmbeddedFileToString(path)
}
