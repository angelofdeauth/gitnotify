// @File:     linux.go
// @Created:  2020-03-23 19:27:52
// @Modified: 2020-03-29 01:10:04
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import "github.com/angelofdeauth/xnotify/pkg/rtc"

// createStartupRscLinux creates a service file for Linux based systems.
func createStartupRscLinux(rtc *rtc.RunTimeCfg) error {

	return createResourceForUser(rtc, &paths{
		templatePath:       "/service/linux-service.gotmpl",
		rootServiceRscPath: "/etc/systemd/user/xnotify.service",
		userServiceRscPath: ".config/systemd/user/xnotify.service",
		rootInstallPath:    "",
		userInstallPath:    "",
	})
}
