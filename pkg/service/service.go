// @File:     service.go
// @Created:  2020-03-21 12:50:33
// @Modified: 2020-03-27 21:19:39
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package service provides an interface for setting up xnotify as a service on the local machine.
package service

import (
	"fmt"
	"runtime"
)

// CreateStartupResources creates the required resources to start the service on boot.
func CreateStartupResources(u string) error {
	switch os := runtime.GOOS; os {
	case "darwin":
		return createStartupRscDarwin(u)
	case "freebsd":
		return createStartupRscFreeBSD(u)
	case "linux":
		return createStartupRscLinux(u)
	case "windows":
		return createStartupRscWindows(u)
	default:
		return fmt.Errorf("Unsupported operating system: %s", os)
	}
}
