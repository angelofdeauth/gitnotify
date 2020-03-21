/*
 * @File:     flg.go
 * @Created:  2020-03-19 14:53:00
 * @Modified: 2020-03-20 00:01:37
 * @Author:   Antonio Escalera
 * @Commiter: Antonio Escalera
 * @Mail:     aj@angelofdeauth.host
 * @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>
 */

package flg

import (
	"fmt"

	"github.com/angelofdeauth/xnotify/pkg/setup"
	"github.com/urfave/cli/v2"
)

// Set sets the cli.App Flags field for global flags.
func Set(a *cli.App, s *setup.Config) {
	var file string
	a.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "file",
			Aliases:     []string{"f"},
			Usage:       "load config from `FILE`",
			Value:       fmt.Sprintf("%s/.config/%s/config.yaml", s.Home, s.App),
			Destination: &file,
		},
	}
}
