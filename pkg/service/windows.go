// @File:     windows.go
// @Created:  2020-03-23 19:30:08
// @Modified: 2020-03-29 01:09:56
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import "github.com/angelofdeauth/xnotify/pkg/rtc"

// createStartupRscWindows creates startup registry keys windows based systems.
func createStartupRscWindows(rtc *rtc.RunTimeCfg) error {

	return createResourceForUser(rtc, &paths{
		templatePath:       "/service/windows-registry.reg.gotmpl",
		rootServiceRscPath: "C:\\Program Files\\xnotify\\xnotify.reg",
		userServiceRscPath: "AppData\\Roaming\\xnotify\\xnotify.reg",
		rootInstallPath:    "",
		userInstallPath:    "",
	})
}
