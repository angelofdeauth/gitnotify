/*
 * @File:     setup.go
 * @Created:  2020-03-19 09:58:40
 * @Modified: 2020-03-19 09:59:46
 * @Author:   Antonio Escalera
 * @Commiter: Antonio Escalera
 * @Mail:     aj@angelofdeauth.host
 * @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>
 */

package xnotify

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type setupConfig struct {
	home string
}

// setupCommon sets up the general user-specific configuration parameters
func (s *setupConfig) setupCommon() {
	usrhome, errh := os.UserHomeDir()
	if errh != nil {
		log.Fatal(errh)
	}
	s.home = usrhome
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
	s.setupCommon()

	return &s
}
