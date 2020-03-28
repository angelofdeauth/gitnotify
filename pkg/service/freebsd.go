// @File:     freebsd.go
// @Created:  2020-03-23 19:28:59
// @Modified: 2020-03-28 01:18:52
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createStartupRscFreeBSD creates a service file for FreeBSD based systems.
func createStartupRscFreeBSD(u string) error {

	return createResourceForUser(u, "/service/freebsd-service.gotmpl",
		"/etc/rc.d/xnotify",
		"/usr/local/etc/rc.d/xnotify.sh")
}
