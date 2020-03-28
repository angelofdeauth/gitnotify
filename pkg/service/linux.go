// @File:     linux.go
// @Created:  2020-03-23 19:27:52
// @Modified: 2020-03-28 01:18:38
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createStartupRscLinux creates a service file for Linux based systems.
func createStartupRscLinux(u string) error {

	return createResourceForUser(u, "/service/linux-service.gotmpl",
		"/etc/systemd/user/xnotify.service",
		".config/systemd/user/xnotify.service")
}
