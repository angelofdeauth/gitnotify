/*
 * @File:     main.go
 * @Created:  2020-03-18 21:47:51
 * @Modified: 2020-03-19 16:32:10
 * @Author:   Antonio Escalera
 * @Commiter: Antonio Escalera
 * @Mail:     aj@angelofdeauth.host
 * @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>
 */

/* `main.go` is the entrypoint for the application.
 *
 *
 */

package main

import (
    "fmt"
    "log"
    "os"
    "sort"
    "time"

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

func main() {
    s := setup.Init()

    app := &cli.App{
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

    sort.Sort(cli.FlagsByName(app.Flags))
    sort.Sort(cli.CommandsByName(app.Commands))

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
