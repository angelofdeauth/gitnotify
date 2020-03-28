// @File:     freebsd.go
// @Created:  2020-03-23 19:28:59
// @Modified: 2020-03-28 03:59:59
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createStartupRscFreeBSD creates a service file for FreeBSD based systems.
func (sf *Flags) createStartupRscFreeBSD() error {

	return sf.createResourceForUser("/service/freebsd-service.gotmpl",
		"/etc/rc.d/xnotify",
		"/usr/local/etc/rc.d/xnotify.sh")
}
