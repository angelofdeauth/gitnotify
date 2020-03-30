// @File:     darwin.go
// @Created:  2020-03-23 18:30:06
// @Modified: 2020-03-29 23:58:37
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

// launchctlLoad is the function to load the plist into launchctl.
func launchctlLoad(srv string) ([]byte, error) {
	return exec.Command("/bin/launchctl", "load", "-w", srv).CombinedOutput()
}

// startDaemonDarwin applies the
func startDaemonDarwin(rtc *rtc.RunTimeCfg) error {

	o, err := launchctlLoad(rtc.OutputPath)
	if err != nil {
		return err
	}
	log.Printf("[launchctl output]: %s", o)

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
	if rtc.Start {
		if err := startDaemonDarwin(rtc); err != nil {
			return err
		}
	}
	return nil
}
