// @File:     write.go
// @Created:  2020-03-23 15:25:31
// @Modified: 2020-03-29 01:29:20
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import (
  "github.com/angelofdeauth/xnotify/pkg/box"
  "github.com/angelofdeauth/xnotify/pkg/rtc"
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
func createResourceForUser(rtc *rtc.RunTimeCfg, pths *paths) error {

  // set up fileAttributes
  fa := box.NewFileAttributes()
  fa.TemplatePath = pths.templatePath

  // switch behavior based on status of the `-u/--user` flag.
  if rtc.User == "root" {
    // default option, no user flag specified.

    // set the config object install path for use in templates
    rtc.InstallPath = rtc.GetInstallPathD(pths.rootInstallPath)

    //set the runtime config output path
    rtc.OutputPath = rtc.GetOutputPathD(pths.rootServiceRscPath)

    // set the file attributes output path
    fa.OutputPath = rtc.OutputPath

    // the defaults from `box.NewFileAttributes()` are used for the rest of the file attribute fields.

  } else {
    // user specified from flag or unspecified user flag set.

    // set the config object install path for use in templates
    rtc.InstallPath = rtc.GetInstallPathD(pths.userInstallPath)

    // set the runtime config output path
    rtc.OutputPath = rtc.GetOutputPathD(pths.userServiceRscPath)

    // set the file attributes for passed user.
    err := fa.SetFileAttributesForUser(rtc.User, rtc.OutputPath)
    if err != nil {
      return err
    }
  }

  // send set up variables to be rendered through fileFromTemplate
  if err := box.Template.WriteFileFromTemplate(fa, rtc); err != nil {
    return err
  }

  return nil
}
