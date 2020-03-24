// @File:     cmd.go
// @Created:  2020-03-19 19:05:29
// @Modified: 2020-03-23 15:46:40
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package cmd sets up the Command field.
package cmd

import (
	"github.com/angelofdeauth/xnotify/pkg/service"
	"github.com/urfave/cli/v2"
)

// Set sets the cli.App Commands field.
func Set(a *cli.App) {
	var serviceUser string

	a.Commands = []*cli.Command{
		{
			Name:    "daemon",
			Aliases: []string{"d"},
			Usage:   "Start as a daemon.",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Set live and persistent configuration.",
			Action: func(c *cli.Context) error {
				return nil
			},
			Subcommands: []*cli.Command{
				{
					Name:    "add",
					Aliases: []string{"a"},
					Usage:   "Add a directory to config.",
					Action: func(c *cli.Context) error {
						return nil
					},
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:    "parents",
							Aliases: []string{"p"},
							Usage:   "no error if existing, make parent directories as needed",
							Value:   true,
						},
					},
				},
				{
					Name:    "remove",
					Aliases: []string{"r"},
					Usage:   "Remove a directory from config.",
					Action: func(c *cli.Context) error {
						return nil
					},
				},
			},
		},
		{
			Name:    "service",
			Aliases: []string{"s"},
			Usage:   "Set up service file.",
			Action: func(c *cli.Context) error {
				return service.CreateStartupResources(serviceUser)
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "user",
					Aliases:     []string{"u"},
					Usage:       "run service as `USER`",
					Value:       "root",
					Destination: &serviceUser,
				},
			},
		},
	}
}
