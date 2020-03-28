// @File:     darwin.go
// @Created:  2020-03-23 18:30:06
// @Modified: 2020-03-28 15:59:47
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createStartupRscDarwin creates a launchd job file for macOS.
func (sf *Config) createStartupRscDarwin() error {

	return sf.createResourceForUser(&paths{
		templatePath:       "/service/darwin-launchd.plist.gotmpl",
		rootServiceRscPath: "/Library/LaunchDaemons/xnotify.plist",
		userServiceRscPath: "Library/LaunchAgents/com.angelofdeauth.fate.sdd.xnotify.plist",
		rootInstallPath:    "",
		userInstallPath:    "",
	})
}
