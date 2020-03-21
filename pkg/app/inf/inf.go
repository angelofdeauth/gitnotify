/*
 * @File:     inf.go
 * @Created:  2020-03-19 18:38:41
 * @Modified: 2020-03-19 20:41:04
 * @Author:   Antonio Escalera
 * @Commiter: Antonio Escalera
 * @Mail:     aj@angelofdeauth.host
 * @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>
 */

package inf

import (
	"time"

	"github.com/angelofdeauth/xnotify/pkg/setup"
	"github.com/urfave/cli/v2"
)

// Set sets the information fields of the cli.App.
func Set(a *cli.App, s *setup.Config) {
	a.Name = s.App
	a.Version = s.Version
	a.Compiled = time.Now().UTC()
	a.Copyright = "Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>"
	a.Description = "xnotify is a filesystem event based workflow automation daemon.\n"
	a.Usage = "A filesystem event based workflow automation daemon."
	a.Metadata = map[string]interface{}{
		"Commit": s.Commit,
	}
}
