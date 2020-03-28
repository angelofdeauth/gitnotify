// @File:     windows.go
// @Created:  2020-03-23 19:30:08
// @Modified: 2020-03-28 15:58:34
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createStartupRscWindows creates startup registry keys windows based systems.
func (sf *Config) createStartupRscWindows() error {

	return sf.createResourceForUser(&paths{
		templatePath:       "/service/windows-registry.reg.gotmpl",
		rootServiceRscPath: "C:\\Program Files\\xnotify\\xnotify.reg",
		userServiceRscPath: "AppData\\Roaming\\xnotify\\xnotify.reg",
		rootInstallPath:    "",
		userInstallPath:    "",
	})
}
