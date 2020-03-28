// @File:     windows.go
// @Created:  2020-03-23 19:30:08
// @Modified: 2020-03-28 01:17:55
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createStartupRscWindows creates startup registry keys windows based systems.
func createStartupRscWindows(u string) error {
  return createResourceForUser(u, "/service/windows-registry.reg.gotmpl",
    "C:\\Program Files\\xnotify\\xnotify.reg",
    "AppData\\Roaming\\xnotify\\xnotify.reg")
}
