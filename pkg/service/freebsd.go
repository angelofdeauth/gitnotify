// @File:     freebsd.go
// @Created:  2020-03-23 19:28:59
// @Modified: 2020-03-29 12:46:21
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import "github.com/angelofdeauth/xnotify/pkg/rtc"

// applyStartupRscFreeBSD applies the
func applyStartupRscFreeBSD(rtc *rtc.RunTimeCfg) error {
	return nil
}

// createStartupRscFreeBSD creates a service file for FreeBSD based systems.
func createStartupRscFreeBSD(rtc *rtc.RunTimeCfg) error {

	return writeResourceForUser(rtc, &paths{
		templatePath:       "/service/freebsd-service.gotmpl",
		rootServiceRscPath: "/etc/rc.d/xnotify",
		userServiceRscPath: "/usr/local/etc/rc.d/xnotify.sh",
		rootInstallPath:    "",
		userInstallPath:    "",
	})
}

// FreeBSD is the service setup function called for FreeBSD based systems.
func FreeBSD(rtc *rtc.RunTimeCfg) error {
	if err := createStartupRscFreeBSD(rtc); err != nil {
		return err
	}
	if rtc.Apply {
		if err := applyStartupRscFreeBSD(rtc); err != nil {
			return err
		}
	}
	return nil
}
