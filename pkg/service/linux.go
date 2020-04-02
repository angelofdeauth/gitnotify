// @File:     linux.go
// @Created:  2020-03-23 19:27:52
// @Modified: 2020-03-30 23:40:26
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

// systemctlStartCmd is the function to load and start the service.
func systemctlStartCmd(srv string) error {

	strdr, err := exec.Command("systemctl", "daemon-reload").CombinedOutput()
	if err != nil {
		return err
	}
	log.Printf("%s", strdr)

	stren, err := exec.Command("systemctl", "enable", srv).CombinedOutput()
	if err != nil {
		return err
	}
	log.Printf("%s", stren)

	strst, err := exec.Command("systemctl", "start", srv).CombinedOutput()
	if err != nil {
		return err
	}
	log.Printf("%s", strst)

	return nil
}

// startDaemonLinux applies the
func startDaemonLinux(rtc *rtc.RunTimeCfg) error {

	if err := systemctlStartCmd(rtc.App); err != nil {
		return err
	}

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
