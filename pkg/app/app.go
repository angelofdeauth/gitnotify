// @File:     app.go
// @Created:  2020-03-19 20:02:09
// @Modified: 2020-03-21 04:06:09
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package app sets up the cli app through the urfave/cli/v2 library.
// This package provides a more manageable interface when dealing with large projects.
package app

import (
    "github.com/angelofdeauth/xnotify/pkg/app/aut"
    "github.com/angelofdeauth/xnotify/pkg/app/cmd"
    "github.com/angelofdeauth/xnotify/pkg/app/flg"
    "github.com/angelofdeauth/xnotify/pkg/app/inf"
    "github.com/angelofdeauth/xnotify/pkg/app/opt"
    "github.com/angelofdeauth/xnotify/pkg/setup"
    "github.com/urfave/cli/v2"
)

// Init initializes the app and runs the functions to set the info, authors, commands, and flags.
func Init(s *setup.Config) (*cli.App, error) {
    a := &cli.App{}
    inf.Set(a, s)
    opt.Set(a)
    aut.Set(a)
    cmd.Set(a)
    flg.Set(a, s)
    return a, nil
}
