// @File:     write.go
// @Created:  2020-03-23 15:25:31
// @Modified: 2020-03-28 16:14:32
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import (
  "github.com/angelofdeauth/xnotify/pkg/box"
)

type paths struct {
  rootInstallPath    string
  rootServiceRscPath string
  templatePath       string
  userInstallPath    string
  userServiceRscPath string
}

// createResourcesForUser generic function for creating resource files.
// inputs: username (string), template path in box (string), root destination path (string), user destination path (string)
func (sf *Config) createResourceForUser(pths *paths) error {

  // set up fileAttributes
  fa := box.NewFileAttributes()
  fa.TemplatePath = pths.templatePath

  // switch behavior based on status of the `-u/--user` flag.
  if sf.User == "root" {
    // default option, no user flag specified.

    // set the config object install path for use in templates
    sf.InstallPath = sf.getInstallPathD(pths.rootInstallPath)

    // set the file attributes output path
    fa.OutputPath = sf.getOutputPathD(pths.rootServiceRscPath)

    // the defaults from `box.NewFileAttributes()` are used for the rest of the file attribute fields.

  } else {
    // user specified from flag or unspecified user flag set.

    // set the config object install path for use in templates
    sf.InstallPath = sf.getInstallPathD(pths.userInstallPath)

    // set the file attributes for passed user.
    err := fa.SetFileAttributesForUser(sf.User, sf.getOutputPathD(pths.userServiceRscPath))
    if err != nil {
      return err
    }
  }

  // send set up variables to be rendered through fileFromTemplate
  if err := box.Template.WriteFileFromTemplate(fa, sf); err != nil {
    return err
  }

  return nil
}
