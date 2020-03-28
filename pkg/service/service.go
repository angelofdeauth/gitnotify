// @File:     service.go
// @Created:  2020-03-21 12:50:33
// @Modified: 2020-03-28 04:27:59
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

// Flags is the object for the flags passed to the service command.
type Flags struct {
	InstallDir string
	OutputDir  string
	User       string
}

// getInstallDirDefault returns the Flags.InstallDir field if set, or the passed string otherwise.
func (sf *Flags) getInstallDirDefault(str string) string {

	if sf.InstallDir == "" {
		return str
	}
	return sf.InstallDir
}

// getOutputDirDefault returns the Flags.OutputDir field if set, or the passed string otherwise.
func (sf *Flags) getOutputDirDefault(str string) string {

	if sf.OutputDir == "" {
		return str
	}
	return sf.OutputDir
}

// CreateStartupResources creates the required resources to start the service on boot.
func (sf *Flags) CreateStartupResources() error {
	switch os := runtime.GOOS; os {
	case "darwin":
		return sf.createStartupRscDarwin()
	case "freebsd":
		return sf.createStartupRscFreeBSD()
	case "linux":
		return sf.createStartupRscLinux()
	case "windows":
		return sf.createStartupRscWindows()
	default:
		return fmt.Errorf("Unsupported operating system: %s", os)
	}
}
