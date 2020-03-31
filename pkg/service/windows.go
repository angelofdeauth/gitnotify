// @File:     windows.go
// @Created:  2020-03-23 19:30:08
// @Modified: 2020-03-31 00:36:44
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

import (
	"log"
	"os/exec"

	"github.com/angelofdeauth/xnotify/pkg/rtc"
)

// netStartService is the function to start the service.
func netStartService(srv string) ([]byte, error) {
	return exec.Command("net", "start", srv).CombinedOutput()
}

// startDaemonWindows applies the
func startDaemonWindows(rtc *rtc.RunTimeCfg) error {
	o, err := netStartService(rtc.App)
	if err != nil {
		return err
	}
	log.Printf("[net start output]: %s", o)

	return nil
}

// applyStartupRscWindows applies the registry keys generated from createStartupRscWindows.
func applyStartupRscWindows(rtc *rtc.RunTimeCfg) ([]byte, error) {
	return exec.Command("regedit", "\\s", rtc.OutputPath).CombinedOutput()
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
	strap, err := applyStartupRscWindows(rtc)
	if err != nil {
		return err
	}
	log.Printf("[regedit output]: %s", strap)
	if rtc.Start {
		if err := startDaemonWindows(rtc); err != nil {
			return err
		}
	}
	return nil
}
