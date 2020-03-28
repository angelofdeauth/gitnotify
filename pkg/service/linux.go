// @File:     linux.go
// @Created:  2020-03-23 19:27:52
// @Modified: 2020-03-28 15:59:39
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createStartupRscLinux creates a service file for Linux based systems.
func (sf *Config) createStartupRscLinux() error {

	return sf.createResourceForUser(&paths{
		templatePath:       "/service/linux-service.gotmpl",
		rootServiceRscPath: "/etc/systemd/user/xnotify.service",
		userServiceRscPath: ".config/systemd/user/xnotify.service",
		rootInstallPath:    "",
		userInstallPath:    "",
	})
}
