// @File:     freebsd.go
// @Created:  2020-03-23 19:28:59
// @Modified: 2020-03-27 17:28:05
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createServiceFileFreeBSD creates a service file for FreeBSD based systems.
func createServiceFileFreeBSD(u string) error {

	return createResourceForUser(u, "/service/freebsd-service.gotmpl", "", "")
}
