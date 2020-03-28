// @File:     write.go
// @Created:  2020-03-23 15:25:31
// @Modified: 2020-03-28 04:37:43
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import (
  "github.com/angelofdeauth/xnotify/pkg/box"
)

// createResourcesForUser generic function for creating resource files.
// inputs: username (string), template path in box (string), root destination path (string), user destination path (string)
func (sf *Flags) createResourceForUser(tpth string, rpth string, upth string) error {

  // set up fileAttributes
  fa := box.NewFileAttributes()
  fa.TemplatePath = tpth

  // switch behavior based on status of the `-u/--user` flag.
  if sf.User == "root" {

    // default option, no user flag specified.
    fa.OutputPath = sf.getOutputDirDefault(rpth)

    // the defaults from `box.NewFileAttributes()` are used for the rest of the file attribute fields.

  } else {

    // user specified from flag or unspecified user flag set.
    err := fa.SetFileAttributesForUser(sf.User, sf.getOutputDirDefault(upth))
    if err != nil {
      return err
    }
  }

  // send set up variables to be rendered through fileFromTemplate
  if err := box.Template.WriteFileFromTemplate(fa, nil); err != nil {
    return err
  }

  return nil
}
