// @File:     windows.go
// @Created:  2020-03-23 19:30:08
// @Modified: 2020-03-28 03:59:31
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createStartupRscWindows creates startup registry keys windows based systems.
func (sf *Flags) createStartupRscWindows() error {

	return sf.createResourceForUser("/service/windows-registry.reg.gotmpl",
		"C:\\Program Files\\xnotify\\xnotify.reg",
		"AppData\\Roaming\\xnotify\\xnotify.reg")
}
