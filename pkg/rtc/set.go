// @File:     set.go
// @Created:  2020-03-29 00:55:11
// @Modified: 2020-03-29 23:03:11
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package rtc

import "os"

// SetCurrentUserHome sets the current user's home directory
func (rtc *RunTimeCfg) SetCurrentUserHome() error {

	// set run-time user's home directory
	var err error
	if rtc.Home, err = os.UserHomeDir(); err != nil {
		return err
	}

	return nil
}

// SetInstallPath sets the install path to the given `str` input.
func (rtc *RunTimeCfg) SetInstallPath(str string) {
	rtc.InstallPath = str
}

// SetOutputPath sets the output path to the given `str` input
func (rtc *RunTimeCfg) SetOutputPath(str string) {
	rtc.OutputPath = str
}
