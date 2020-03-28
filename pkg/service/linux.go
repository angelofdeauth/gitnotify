// @File:     linux.go
// @Created:  2020-03-23 19:27:52
// @Modified: 2020-03-27 17:29:57
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createServiceFileLinux creates a service file for Linux based systems.
func createServiceFileLinux(u string) error {

	return createResourceForUser(u, "/service/linux-service.gotmpl", "", "")
}
