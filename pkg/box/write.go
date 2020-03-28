// @File:     write.go
// @Created:  2020-03-27 16:04:02
// @Modified: 2020-03-27 22:10:43
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
	if _, err := os.Stat(d); os.IsNotExist(err) {
		if err := os.MkdirAll(d, 0750); err != nil {
			return err
		}
	}

	// create file
	file, err := os.Create(fa.OutputPath)
	if err != nil {
		return err
	}

	// read embedded template to string
	tmpl, err := e.ReadEmbeddedTemplateToString(fa.TemplatePath)
	if err != nil {
		return err
	}

	// create template from loaded string
	t := template.Must(template.New("").Parse(tmpl))

	// render template to file
	if err := t.Execute(file, config); err != nil {
		return err
	}

	// change mode of file
	if err := os.Chmod(fa.OutputPath, fa.Mode); err != nil {
		return err
	}

	// change ownership of file
	if err := os.Chown(fa.OutputPath, fa.Owners.UID, fa.Owners.GID); err != nil {
		return err
	}

	// reset the embed decoder
	if err := EmbedDecoder.Reset(nil); err != nil {
		return err
	}

	return file.Close()
}
