// @File:     darwin.go
// @Created:  2020-03-23 18:30:06
// @Modified: 2020-03-24 18:56:44
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package service

// createLaunchdJobDarwin creates a launchd job file for macOS.
func createLaunchdJobDarwin(u string) error {
	/*
		// set up launchd job file template variable
		template, err := box.Template.ReadEmbeddedTemplateToString("/service/darwin-launchd.plist.gotmpl")
		if err != nil {
			if errr := box.EmbedDecoder.Reset(nil); errr != nil {
				return errr
			}
			return err
		}

		// set up output path, mode, and owners variables
		path := ""
		mode := 0770
		fowners := owners{}
		switch u {
		case "root":
			path = "/Library/LaunchDaemons/xnotify.plist"
			fowners.uid = 0
			fowners.gid = 0
		case "":
			path = "~/Library/LaunchAgents"
			usr, err := user.Current()
			if err != nil {
				return err
			}
			fowners.uid, err = strconv.Atoi(usr.Uid)
			fowners.uid, err = strconv.Atoi(usr.Gid)
		default:
			path = "~/Library/LaunchAgents"
			usr, err := user.Lookup(u)
			if err != nil {
				return err
			}

			fowners.uid, err = strconv.Atoi(usr.Uid)
			if err != nil {

			}
			fowners.uid, err = strconv.Atoi(usr.Gid)
		}

		// send set up variables to be rendered through fileFromTemplate
		if err := box.Template.WriteFileFromTemplate(template, path, mode, owners, interface{}); err != nil {
			if errr := box.EmbedDecoder.Reset(nil); errr != nil {
				return errr
			}
			return err
		}

		if err := box.EmbedDecoder.Reset(nil); err != nil {
			return err
		}*/
	return nil
}
