// @File:     darwin.go
// @Created:  2020-03-23 18:30:06
// @Modified: 2020-03-29 12:46:15
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import "github.com/angelofdeauth/xnotify/pkg/rtc"

// applyStartupRscDarwin applies the
func applyStartupRscDarwin(rtc *rtc.RunTimeCfg) error {
	return nil
}

// createStartupRscDarwin creates a launchd job file for macOS.
func createStartupRscDarwin(rtc *rtc.RunTimeCfg) error {

	if err := writeResourceForUser(rtc, &paths{
		templatePath:       "/service/darwin-launchd.plist.gotmpl",
		rootServiceRscPath: "/Library/LaunchDaemons/xnotify.plist",
		userServiceRscPath: "Library/LaunchAgents/com.angelofdeauth.fate.sdd.xnotify.plist",
		rootInstallPath:    "",
		userInstallPath:    "",
	}); err != nil {
		return err
	}

	return nil
}

// Darwin is service setup function called for macOS based systems.
func Darwin(rtc *rtc.RunTimeCfg) error {
	if err := createStartupRscDarwin(rtc); err != nil {
		return err
	}
	if rtc.Apply {
		if err := applyStartupRscDarwin(rtc); err != nil {
			return err
		}
	}
	return nil
}
