// @File:     app.go
// @Created:  2020-03-19 20:02:09
// @Modified: 2020-03-29 01:31:42
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package app sets up the cli app through the urfave/cli/v2 library.
// This package provides a more manageable interface when dealing with large projects.
package app

import (
  "sort"

  "github.com/angelofdeauth/xnotify/pkg/app/aut"
  "github.com/angelofdeauth/xnotify/pkg/app/cmd"
  "github.com/angelofdeauth/xnotify/pkg/app/flg"
  "github.com/angelofdeauth/xnotify/pkg/app/inf"
  "github.com/angelofdeauth/xnotify/pkg/app/opt"
  "github.com/angelofdeauth/xnotify/pkg/rtc"
  "github.com/urfave/cli/v2"
)

var (

  // APP is the name of the application
  APP = ""

  // COMMIT is the commit hash
  COMMIT = ""

  // TIME is the time of compilation
  TIME = ""

  // VERSION is the package version
  VERSION = ""
)

// Setup initializes the cli.App and runs the functions to set the info, authors, commands, and flags.
func Setup() (*cli.App, error) {

  RTC, err := rtc.NewRunTimeCfg(APP, COMMIT, TIME, VERSION)

  if overrides() != nil {
    return nil, err
  }

  a := &cli.App{}
  inf.Set(a, RTC)
  opt.Set(a)
  aut.Set(a)
  cmd.Set(a, RTC)
  flg.Set(a, RTC)
  sort.Sort(cli.FlagsByName(a.Flags))
  sort.Sort(cli.CommandsByName(a.Commands))
  return a, nil
}
