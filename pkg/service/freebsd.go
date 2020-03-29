// @File:     freebsd.go
// @Created:  2020-03-23 19:28:59
// @Modified: 2020-03-29 01:10:09
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import "github.com/angelofdeauth/xnotify/pkg/rtc"

// createStartupRscFreeBSD creates a service file for FreeBSD based systems.
func createStartupRscFreeBSD(rtc *rtc.RunTimeCfg) error {

	return createResourceForUser(rtc, &paths{
		templatePath:       "/service/freebsd-service.gotmpl",
		rootServiceRscPath: "/etc/rc.d/xnotify",
		userServiceRscPath: "/usr/local/etc/rc.d/xnotify.sh",
		rootInstallPath:    "",
		userInstallPath:    "",
	})
}
