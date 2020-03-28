// @File:     freebsd.go
// @Created:  2020-03-23 19:28:59
// @Modified: 2020-03-28 15:59:43
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createStartupRscFreeBSD creates a service file for FreeBSD based systems.
func (sf *Config) createStartupRscFreeBSD() error {

	return sf.createResourceForUser(&paths{
		templatePath:       "/service/freebsd-service.gotmpl",
		rootServiceRscPath: "/etc/rc.d/xnotify",
		userServiceRscPath: "/usr/local/etc/rc.d/xnotify.sh",
		rootInstallPath:    "",
		userInstallPath:    "",
	})
}
