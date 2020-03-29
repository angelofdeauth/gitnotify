// @File:     flg.go
// @Created:  2020-03-20 22:28:12
// @Modified: 2020-03-29 00:23:18
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package flg sets up the Flag field.
package flg

import (
	"fmt"

	"github.com/angelofdeauth/xnotify/pkg/rtc"
	"github.com/urfave/cli/v2"
)

// Set sets the cli.App Flags field for global flags.
func Set(a *cli.App, rtc *rtc.RunTimeCfg) {

	a.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "file",
			Aliases:     []string{"f"},
			Usage:       "load config from `FILE`",
			Value:       fmt.Sprintf("%s/.config/%s/config.yaml", rtc.Home, rtc.App),
			Destination: &rtc.File,
		},
	}
}
