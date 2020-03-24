// @File:     windows.go
// @Created:  2020-03-23 19:30:08
// @Modified: 2020-03-23 19:30:14
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import "github.com/angelofdeauth/xnotify/pkg/box"

// createRegistryKeysWindows creates startup registry keys windows based systems.
func createRegistryKeysWindows(u string) error {
	rkw, err := box.TemplateBox.ReadEmbeddedTemplateToString("/service/windows-registry.reg.gotmpl")
	if err != nil {
		return err
	}

}
