// @File:     freebsd.go
// @Created:  2020-03-23 19:28:59
// @Modified: 2020-03-23 19:29:41
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import (
	"github.com/angelofdeauth/xnotify/pkg/box"
)

// createServiceFileFreeBSD creates a service file for FreeBSD based systems.
func createServiceFileFreeBSD(u string) error {

	// set up service file template variable
	template, err := box.TemplateBox.ReadEmbeddedTemplateToString("/service/freebsd-service.gotmpl")
	if err != nil {
		return err
	}

}
