// @File:     service.go
// @Created:  2020-03-21 12:50:33
// @Modified: 2020-03-29 00:56:34
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package service provides an interface for setting up xnotify as a service on the local machine.
package service

import (
	"fmt"
	"runtime"

	"github.com/angelofdeauth/xnotify/pkg/rtc"
)

// CreateStartupResources creates the required resources to start the service on boot.
func CreateStartupResources(rtc *rtc.RunTimeCfg) error {
	switch os := runtime.GOOS; os {
	case "darwin":
		return createStartupRscDarwin(rtc)
	case "freebsd":
		return createStartupRscFreeBSD(rtc)
	case "linux":
		return createStartupRscLinux(rtc)
	case "windows":
		return createStartupRscWindows(rtc)
	default:
		return fmt.Errorf("Unsupported operating system: %s", os)
	}
}
