// @File:     freebsd.go
// @Created:  2020-03-23 19:28:59
// @Modified: 2020-03-30 23:57:15
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

// sysrcStartCmd is the function to load and start the service.
func sysrcStartCmd(srv string) error {

	stren, err := exec.Command("sysrc", "xnotify_enable=\"YES\"").CombinedOutput()
	if err != nil {
		return err
	}
	log.Printf("[sysrc ouptut]: %s", stren)

	strst, err := exec.Command("service", srv, "start").CombinedOutput()
	if err != nil {
		return err
	}
	log.Printf("[service output]: %s", strst)

	return nil
}

// startDaemonFreeBSD applies the
func startDaemonFreeBSD(rtc *rtc.RunTimeCfg) error {

	err := sysrcStartCmd(rtc.App)
	if err != nil {
		return err
	}

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
	if rtc.Start {
		if err := startDaemonFreeBSD(rtc); err != nil {
			return err
		}
	}
	return nil
}
