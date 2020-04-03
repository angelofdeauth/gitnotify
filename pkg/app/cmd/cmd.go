// @File:     cmd.go
// @Created:  2020-03-19 19:05:29
// @Modified: 2020-04-03 12:03:55
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package cmd sets up the Command field.
package cmd

import (
	"fmt"
	"runtime"

	"github.com/angelofdeauth/xnotify/pkg/rtc"
	"github.com/angelofdeauth/xnotify/pkg/service"
	"github.com/urfave/cli/v2"
)

// Set sets the cli.App Commands field.
func Set(a *cli.App, rtc *rtc.RunTimeCfg) {

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
			Usage:   "Create service file.",
			Action: func(c *cli.Context) error {
				switch os := runtime.GOOS; os {
				case "darwin":
					return service.Darwin(rtc)
				case "freebsd":
					return service.FreeBSD(rtc)
				case "linux":
					return service.Linux(rtc)
				case "windows":
					return service.Windows(rtc)
				default:
					return fmt.Errorf("Unsupported operating system: %s", os)
				}
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "user",
					Aliases:     []string{"u"},
					Usage:       "run service as `USER`",
					Value:       "root",
					Destination: &rtc.User,
				},
				&cli.StringFlag{
					Name:        "install-path",
					Aliases:     []string{"i"},
					Usage:       "point service file to `FILE`",
					Value:       "",
					Destination: &rtc.InstallPath,
					Required:    true,
				},
				&cli.StringFlag{
					Name:        "output-path",
					Aliases:     []string{"o"},
					Usage:       "output rendered service resources to `DIR`",
					Value:       "",
					Destination: &rtc.OutputPath,
				},
				&cli.BoolFlag{
					Name:        "start",
					Aliases:     []string{"s"},
					Value:       false,
					Destination: &rtc.Start,
				},
				&cli.StringFlag{
					Name:        "daemon-args",
					Aliases:     []string{"d"},
					Value:       "",
					Destination: &rtc.DaemonArgs,
				},
			},
		},
	}
}
