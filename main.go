/*
 * @File:     main.go
 * @Created:  2020-03-18 21:47:51
 * @Modified: 2020-03-19 20:05:05
 * @Author:   Antonio Escalera
 * @Commiter: Antonio Escalera
 * @Mail:     aj@angelofdeauth.host
 * @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>
 */

// Package main is the entry point of the application.
package main

import (
    "log"
    "os"
    "sort"

    "github.com/angelofdeauth/xnotify/pkg/app"
    "github.com/angelofdeauth/xnotify/pkg/setup"
    "github.com/urfave/cli/v2"
)

var (

    // APP is the name of the application
    APP = ""

    // COMMIT is the commit hash
    COMMIT = ""

    // VERSION is the package version
    VERSION = ""
)

func exit(e error) {
    log.Fatal(e)
}

func main() {
    s, errs := setup.Init(APP, COMMIT, VERSION)
    if errs != nil {
        exit(errs)
    }

    a, errb := app.Init(s)
    if errb != nil {
        exit(errb)
    }

    sort.Sort(cli.FlagsByName(a.Flags))
    sort.Sort(cli.CommandsByName(a.Commands))

    erra := a.Run(os.Args)
    if erra != nil {
        exit(erra)
    }
}
