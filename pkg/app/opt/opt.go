// @File:     opt.go
// @Created:  2020-03-19 18:51:14
// @Modified: 2020-03-21 04:20:07
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package opt sets up various option fields.
package opt

import "github.com/urfave/cli/v2"

// Set sets the cli.App fields for various app behavior options.
func Set(a *cli.App) {
	a.EnableBashCompletion = true
	a.HideHelpCommand = true
	a.UseShortOptionHandling = true
}
