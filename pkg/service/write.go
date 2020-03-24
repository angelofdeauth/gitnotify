// @File:     write.go
// @Created:  2020-03-23 15:25:31
// @Modified: 2020-03-24 02:58:11
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import (
	"os"
	"path/filepath"
	"text/template"
)

// owners is the file ownership structure.
type owners struct {
	uid int
	gid int
}

// WriteFileFromTemplate creates startup resources from given templates.
func WriteFileFromTemplate(input string, outpath string, outmod os.FileMode, outown owners, config interface{}) error {

	// create directory (if required)
	d := filepath.Dir(outpath)
	if _, errs := os.Stat(d); os.IsNotExist(errs) {
		if err := os.MkdirAll(d, 0750); err != nil {
			return err
		}
	}

	// create file
	file, errc := os.Create(outpath)
	if errc != nil {
		return errc
	}

	// create template from loaded string
	t := template.Must(template.New("input").Parse(input))

	// render template to file
	if errr := t.Execute(file, config); errr != nil {
		return errr
	}

	// change mode of file
	if errm := os.Chmod(outpath, outmod); errm != nil {
		return errm
	}

	// change ownership of file
	if erro := os.Chown(outpath, outown.uid, outown.gid); erro != nil {
		return erro
	}

	return file.Close()
}
