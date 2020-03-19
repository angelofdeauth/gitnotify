/*
 * @File:     setup.go
 * @Created:  2020-03-19 14:47:11
 * @Modified: 2020-03-19 16:50:01
 * @Author:   Antonio Escalera
 * @Commiter: Antonio Escalera
 * @Mail:     aj@angelofdeauth.host
 * @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>
 */

/* The `setup` package sets up and  automatically collected configuration information.
 * This information is collected before the cli.App context is created.
 * It is fed to the `app` package for processing when creating the cli.App context through the `main` package.
 */

package setup

import (
    "fmt"
    "os"

    "github.com/angelofdeauth/xnotify/pkg/read"
    "github.com/urfave/cli/v2"
)

// Config is the struct containing automatically collected configuration information.
type Config struct {
    Home string
}

// common sets up the general user-specific configuration parameters
func (s *Config) common() error {
    usrhome, errh := os.UserHomeDir()
    if errh != nil {
        return errh
    }
    s.Home = usrhome
    return nil
}

// overrides sets up the overrides for various vendored packages
func overrides() error {

    t, errr := read.EmbeddedFileToString("/pkg/tpl/AppHelpTemplate.gotmpl")
    if errr != nil {
        return errr
    }

    // overrides the default cli AppHelpTemplate
    cli.AppHelpTemplate = t

    // overrides the default cli VersionPrinter function
    cli.VersionPrinter = func(c *cli.Context) {
        fmt.Printf("%s\n  SemVer: %s\n  Time:   %s\n  Commit: %s\n\n", c.App.Name, c.App.Version, c.App.Compiled, c.App.Metadata["Commit"])
    }
    return nil

}

// Init sets up the environment for the app
func Init() (*Config, error) {
    s := Config{}

    erro := overrides()
    if erro != nil {
        return nil, erro
    }
    errc := s.common()
    if errc != nil {
        return nil, errc
    }

    return &s, nil
}
