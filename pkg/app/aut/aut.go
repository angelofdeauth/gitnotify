// @File:     aut.go
// @Created:  2020-03-19 18:51:21
// @Modified: 2020-03-21 04:19:18
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package aut sets up the Authors field.
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
