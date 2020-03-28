// @File:     darwin.go
// @Created:  2020-03-23 18:30:06
// @Modified: 2020-03-28 04:00:09
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createStartupRscDarwin creates a launchd job file for macOS.
func (sf *Flags) createStartupRscDarwin() error {

	return sf.createResourceForUser("/service/darwin-launchd.plist.gotmpl",
		"/Library/LaunchDaemons/xnotify.plist",
		"Library/LaunchAgents/com.angelofdeauth.fate.sdd.xnotify.plist")
}
