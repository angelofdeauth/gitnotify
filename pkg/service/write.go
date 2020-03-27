// @File:     write.go
// @Created:  2020-03-23 15:25:31
// @Modified: 2020-03-27 17:58:39
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
func createResourceForUser(u string, tpth string, rpth string, upth string) error {

  // set up fileAttributes
  fa := box.NewFileAttributes()
  fa.TemplatePath = tpth

  // switch behavior based on status of the `-u/--user` flag.
  if u == "root" {

    // default option, no user flag specified.
    fa.OutputPath = rpth

  } else {

    // user specified from flag or unspecified.
    err := fa.SetFileAttributesForUser(u, upth)
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
