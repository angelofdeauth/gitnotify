// @File:     windows.go
// @Created:  2020-03-23 19:30:08
// @Modified: 2020-03-27 17:30:11
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createRegistryKeysWindows creates startup registry keys windows based systems.
func createRegistryKeysWindows(u string) error {
	return createResourceForUser(u, "/service/windows-registry.reg.gotmpl", "", "")
}
