// @File:     linux.go
// @Created:  2020-03-23 19:27:52
// @Modified: 2020-03-23 19:28:24
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import (
	"github.com/angelofdeauth/xnotify/pkg/box"
)

// createServiceFileLinux creates a service file for Linux based systems.
func createServiceFileLinux(u string) error {

	// set up service file template variable
	template, err := box.TemplateBox.ReadEmbeddedTemplateToString("/service/linux-service.gotmpl")
	if err != nil {
		return err
	}

}
