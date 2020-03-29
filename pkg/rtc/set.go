// @File:     set.go
// @Created:  2020-03-29 00:55:11
// @Modified: 2020-03-29 00:55:38
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
