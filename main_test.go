// @File:     main_test.go
// @Created:  2020-03-18 21:48:07
// @Modified: 2020-03-21 04:45:44
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

// Package main is the entrypoint of the test suite.
package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ret := m.Run()
	os.Exit(ret)
}
