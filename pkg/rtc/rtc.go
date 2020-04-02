// @File:     rtc.go
// @Created:  2020-03-28 21:05:29
// @Modified: 2020-03-29 23:54:13
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package rtc automatically collects system information.
// This information is collected before the cli.App context is created.
// It is fed to the `app` package for processing when creating the cli.App context through the `main` package.
package rtc

// RunTimeCfg is the object for the runtime configuration passed to the command.
// This includes flags, arguments, and automatically detected configuration.
type RunTimeCfg struct {
	App     string
	Commit  string
	File    string
	Home    string
	Time    string
	Version string
	Service
}

// Service is the object for runtime configuration passed to the Service command.
type Service struct {
	Start       bool
	InstallPath string
	OutputPath  string
	User        string
}

// NewRunTimeCfg creates a new Run Time Configuration.
func NewRunTimeCfg(a string, c string, t string, v string) (*RunTimeCfg, error) {

	rtc := &RunTimeCfg{App: a, Commit: c, Time: t, Version: v}
	if err := rtc.SetCurrentUserHome(); err != nil {
		return nil, err
	}

	return rtc, nil
}
