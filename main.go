/*
 * @File:     main.go
 * @Created:  2020-03-18 21:47:51
 * @Modified: 2020-03-19 01:58:07
 * @Author:   Antonio Escalera
 * @Commiter: Antonio Escalera
 * @Mail:     aj@angelofdeauth.host
 * @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>
 */

package main

import (
    "fmt"
    "log"
    "os"
    "sort"
    "time"

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

type setupConfig struct {
    home string
}

// setupOverrides sets up the overrides for the cli package
func setupOverrides() {

    // overrides the default cli VersionPrinter function
    cli.VersionPrinter = func(c *cli.Context) {
        fmt.Printf("%s\n  SemVer: %s\n  Time:   %s\n  Commit: %s\n\n", c.App.Name, c.App.Version, c.App.Compiled, COMMIT)
    }

    // overrides the default cli AppHelpTemplate
    cli.AppHelpTemplate = `NAME:
  {{.Name}}{{if .Usage}} - {{.Usage}}{{end}}

USAGE:
  {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}

VERSION:
  SemVer: {{.Version}} 
  Time:   {{.Compiled}} 
  Commit: {{.Metadata.commit}} {{end}}{{end}}{{if .Description}}

DESCRIPTION:
  {{.Description}}{{end}}{{if len .Authors}}

AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
  {{range $index, $author := .Authors}}{{if $index}}
  {{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}

COMMANDS:{{range .VisibleCategories}}{{if .Name}}
  {{.Name}}:{{range .VisibleCommands}}
    {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
  {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}

GLOBAL OPTIONS:
  {{range $index, $option := .VisibleFlags}}{{if $index}}
  {{end}}{{$option}}{{end}}{{end}}{{if .Copyright}}

COPYRIGHT:
  {{.Copyright}}{{end}}

`

}

// setupMain sets up the environment for the app
func setupMain() *setupConfig {
    s := setupConfig{}

    setupOverrides()

    home, errh := os.UserHomeDir()
    if errh != nil {
        log.Fatal(errh)
    }
    s.home = home
    return &s
}

func main() {
    s := setupMain()

    app := &cli.App{
        Name:                   APP,
        Version:                VERSION,
        Compiled:               time.Now().UTC(),
        Copyright:              "Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>",
        Description:            "xnotify is a filesystem event based developer workflow automation daemon.\n",
        Usage:                  "A filesystem event based development automation daemon.",
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
                Value:   fmt.Sprintf("%s/.config/%s/config.yaml", s.home, APP),
            },
        },
        Metadata: map[string]interface{}{
            "commit": COMMIT,
        },
    }

    sort.Sort(cli.FlagsByName(app.Flags))
    sort.Sort(cli.CommandsByName(app.Commands))

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
