// @File:     darwin.go
// @Created:  2020-03-23 18:30:06
// @Modified: 2020-03-28 01:19:02
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createStartupRscDarwin creates a launchd job file for macOS.
func createStartupRscDarwin(u string) error {

	return createResourceForUser(u, "/service/darwin-launchd.plist.gotmpl",
		"/Library/LaunchDaemons/xnotify.plist",
		"Library/LaunchAgents/com.angelofdeauth.fate.sdd.xnotify.plist")
}
