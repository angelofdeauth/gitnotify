// @File:     windows.go
// @Created:  2020-03-23 19:30:08
// @Modified: 2020-03-29 23:58:54
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import "github.com/angelofdeauth/xnotify/pkg/rtc"

// startDaemonWindows applies the
func startDaemonWindows(rtc *rtc.RunTimeCfg) error {
	return nil
}

// createStartupRscWindows creates startup registry keys windows based systems.
func createStartupRscWindows(rtc *rtc.RunTimeCfg) error {

	return writeResourceForUser(rtc, &paths{
		templatePath:       "/service/windows-registry.reg.gotmpl",
		rootServiceRscPath: "C:\\Program Files\\xnotify\\xnotify.reg",
		userServiceRscPath: "AppData\\Roaming\\xnotify\\xnotify.reg",
		rootInstallPath:    "",
		userInstallPath:    "",
	})
}

// Windows is the service setup function called for Windows based systems.
func Windows(rtc *rtc.RunTimeCfg) error {
	if err := createStartupRscWindows(rtc); err != nil {
		return err
	}
	if rtc.Start {
		if err := startDaemonWindows(rtc); err != nil {
			return err
		}
	}
	return nil
}
