/*
 * @File:     app.go
 * @Created:  2020-03-19 14:52:20
 * @Modified: 2020-03-19 20:00:19
 * @Author:   Antonio Escalera
 * @Commiter: Antonio Escalera
 * @Mail:     aj@angelofdeauth.host
 * @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>
 */

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

/*
var app := &cli.App{
    Name:                   APP,
    Version:                VERSION,
    Compiled:               time.Now().UTC(),
    Copyright:              "Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>",
    Description:            "xnotify is a filesystem event based workflow automation daemon.\n",
    Usage:                  "A filesystem event based workflow automation daemon.",
    EnableBashCompletion:   true,
    HideHelpCommand:        true,
    UseShortOptionHandling: true,
    Authors: []*cli.Author{
        {
            Name:  "Antonio Escalera",
            Email: "aj@angelofdeauth.host",
        },
    },
    Commands: []*cli.Command{
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
                return nil
            },
        },
    },
    Flags: []cli.Flag{
        &cli.StringFlag{
            Name:    "configfile",
            Aliases: []string{"c"},
            Usage:   "load config from `FILE`",
            Value:   fmt.Sprintf("%s/.config/%s/config.yaml", s.Home, APP),
        },
    },
    Metadata: map[string]interface{}{
        "Commit": COMMIT,
    },
}
*/
