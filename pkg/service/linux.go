// @File:     linux.go
// @Created:  2020-03-23 19:27:52
// @Modified: 2020-03-28 03:59:41
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createStartupRscLinux creates a service file for Linux based systems.
func (sf *Flags) createStartupRscLinux() error {

	return sf.createResourceForUser("/service/linux-service.gotmpl",
		"/etc/systemd/user/xnotify.service",
		".config/systemd/user/xnotify.service")
}
