// @File:     overrides.go
// @Created:  2020-03-29 00:00:31
// @Modified: 2020-03-29 01:33:25
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package app

import (
	"fmt"

	"github.com/angelofdeauth/xnotify/pkg/box"
	"github.com/urfave/cli/v2"
)

// overrides sets up the overrides for various vendored packages
func overrides() error {

	var err error

	// overrides the default cli AppHelpTemplate
	if cli.AppHelpTemplate, err = box.Template.ReadEmbeddedTemplateToString("/help/AppHelpTemplate.gotmpl"); err != nil {
		return err
	}

	// overrides the default cli CommandHelpTemplate
	if cli.CommandHelpTemplate, err = box.Template.ReadEmbeddedTemplateToString("/help/CommandHelpTemplate.gotmpl"); err != nil {
		return err
	}

	// overrides the default cli SubcommandHelpTemplate
	if cli.SubcommandHelpTemplate, err = box.Template.ReadEmbeddedTemplateToString("/help/SubcommandHelpTemplate.gotmpl"); err != nil {
		return err
	}

	// overrides the default cli VersionPrinter function
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("SemVer: %s\n  Time: %s\nCommit: %s\n\n", c.App.Version, c.App.Compiled, c.App.Metadata["Commit"])
	}

	return nil
}
