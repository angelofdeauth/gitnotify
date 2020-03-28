// @File:     write.go
// @Created:  2020-03-27 16:04:02
// @Modified: 2020-03-27 18:48:43
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package box

import (
	"os"
	"path/filepath"
	"text/template"
)

// functions provided for reading from files embedded inside of the executable.

// WriteFileFromTemplate creates startup resources from given templates.
func (e *EmbedBox) WriteFileFromTemplate(fa *FileAttributes, config interface{}) error {

	// create directory (if required)
	d := filepath.Dir(fa.OutputPath)
	if _, errs := os.Stat(d); os.IsNotExist(errs) {
		if err := os.MkdirAll(d, 0750); err != nil {
			return err
		}
	}

	// create file
	file, errc := os.Create(fa.OutputPath)
	if errc != nil {
		return errc
	}

	tmpl, err := e.ReadEmbeddedTemplateToString(fa.TemplatePath)
	if err != nil {
		return err
	}

	// create template from loaded string
	t := template.Must(template.New("").Parse(tmpl))

	// render template to file
	if errr := t.Execute(file, config); errr != nil {
		return errr
	}

	// change mode of file
	if errm := os.Chmod(fa.OutputPath, fa.Mode); errm != nil {
		return errm
	}

	// change ownership of file
	if erro := os.Chown(fa.OutputPath, fa.Owners.UID, fa.Owners.GID); erro != nil {
		return erro
	}

	EmbedDecoder.Reset(nil)

	return file.Close()
}
