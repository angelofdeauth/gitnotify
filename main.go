// @File:     main.go
// @Created:  2020-03-21 04:21:08
// @Modified: 2020-03-29 01:41:06
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package main is the entry point of the application.
package main

import (
  "log"
  "os"

  "github.com/angelofdeauth/xnotify/pkg/app"
)

func exit(e error) {
  log.Fatal(e)
}

func main() {

  a, err := app.Setup()
  if err != nil {
    exit(err)
  }

  err = a.Run(os.Args)
  if err != nil {
    exit(err)
  }
}
