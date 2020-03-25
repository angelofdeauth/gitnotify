// @File:     setup.go
// @Created:  2020-03-21 03:57:46
// @Modified: 2020-03-25 02:24:03
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package setup automatically collects system information.
// This information is collected before the cli.App context is created.
// It is fed to the `app` package for processing when creating the cli.App context through the `main` package.
package setup

import (
    "fmt"
    "os"

    "github.com/angelofdeauth/xnotify/pkg/box"
    "github.com/urfave/cli/v2"
)

// Config is the struct containing automatically collected configuration information.
type Config struct {
    App     string
    Commit  string
    Home    string
    Version string
}

// common sets up the general user-specific configuration parameters
func (s *Config) common(a string, c string, v string) error {
    usrhome, errh := os.UserHomeDir()
    if errh != nil {
        return errh
    }
    s.Home = usrhome
    s.App = a
    s.Commit = c
    s.Version = v
    return nil
}

// overrides sets up the overrides for various vendored packages
func overrides() error {

    // overrides the default cli AppHelpTemplate
    aht, errr := box.Template.ReadEmbeddedTemplateToString("/help/AppHelpTemplate.gotmpl")
    if errr != nil {
        return errr
    }
    cli.AppHelpTemplate = aht

    // overrides the default cli CommandHelpTemplate
    cht, errr := box.Template.ReadEmbeddedTemplateToString("/help/CommandHelpTemplate.gotmpl")
    if errr != nil {
        return errr
    }
    cli.CommandHelpTemplate = cht

    // overrides the default cli SubcommandHelpTemplate
    sht, errr := box.Template.ReadEmbeddedTemplateToString("/help/SubcommandHelpTemplate.gotmpl")
    if errr != nil {
        return errr
    }
    cli.SubcommandHelpTemplate = sht

    // overrides the default cli VersionPrinter function
    cli.VersionPrinter = func(c *cli.Context) {
        fmt.Printf("SemVer: %s\n  Time: %s\nCommit: %s\n\n", c.App.Version, c.App.Compiled, c.App.Metadata["Commit"])
    }
    return nil

}

// Init sets up the environment for the app
func Init(a string, c string, v string) (*Config, error) {
    s := Config{}

    erro := overrides()
    if erro != nil {
        return nil, erro
    }
    errc := s.common(a, c, v)
    if errc != nil {
        return nil, errc
    }

    return &s, nil
}
