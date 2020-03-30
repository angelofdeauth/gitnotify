// @File:     linux.go
// @Created:  2020-03-23 19:27:52
// @Modified: 2020-03-29 23:58:50
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import "github.com/angelofdeauth/xnotify/pkg/rtc"

// startDaemonLinux applies the
func startDaemonLinux(rtc *rtc.RunTimeCfg) error {
	return nil
}

// createStartupRscLinux creates a service file for Linux based systems.
func createStartupRscLinux(rtc *rtc.RunTimeCfg) error {

	return writeResourceForUser(rtc, &paths{
		templatePath:       "/service/linux-service.gotmpl",
		rootServiceRscPath: "/etc/systemd/user/xnotify.service",
		userServiceRscPath: ".config/systemd/user/xnotify.service",
		rootInstallPath:    "",
		userInstallPath:    "",
	})
}

// Linux is the service setup function called for Linux based systems.
func Linux(rtc *rtc.RunTimeCfg) error {
	if err := createStartupRscLinux(rtc); err != nil {
		return err
	}
	if rtc.Start {
		if err := startDaemonLinux(rtc); err != nil {
			return err
		}
	}
	return nil
}
