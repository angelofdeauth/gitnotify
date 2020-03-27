// @File:     darwin.go
// @Created:  2020-03-23 18:30:06
// @Modified: 2020-03-27 17:27:59
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createLaunchdJobDarwin creates a launchd job file for macOS.
func createLaunchdJobDarwin(u string) error {

	return createResourceForUser(u, "/service/darwin-launchd.plist.gotmpl", "/Library/LaunchDaemons/", "Library/LaunchAgents")
}
