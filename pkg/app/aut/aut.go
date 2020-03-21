/*
 * @File:     aut.go
 * @Created:  2020-03-19 18:48:14
 * @Modified: 2020-03-19 18:51:20
 * @Author:   Antonio Escalera
 * @Commiter: Antonio Escalera
 * @Mail:     aj@angelofdeauth.host
 * @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>
 */

package aut

import "github.com/urfave/cli/v2"

// Set sets the cli.App Authors field.
func Set(a *cli.App) {
	a.Authors = []*cli.Author{
		{
			Name:  "Antonio Escalera",
			Email: "aj@angelofdeauth.host",
		},
	}
}
