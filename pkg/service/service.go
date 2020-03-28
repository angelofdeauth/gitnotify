// @File:     service.go
// @Created:  2020-03-21 12:50:33
// @Modified: 2020-03-28 16:17:19
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

// Config is the object for the flags passed to the service command.
type Config struct {
	App         string
	InstallPath string
	OutputPath  string
	User        string
}

// getInstallPathD returns the Config.InstallPath field if set, or the passed string otherwise.
func (sf *Config) getInstallPathD(str string) string {

	if sf.InstallPath == "" {
		return str
	}
	return sf.InstallPath
}

// getOutputPathD returns the Config.OutputPath field if set, or the passed string otherwise.
func (sf *Config) getOutputPathD(str string) string {

	if sf.OutputPath == "" {
		return str
	}
	return sf.OutputPath
}

// CreateStartupResources creates the required resources to start the service on boot.
func (sf *Config) CreateStartupResources() error {
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
