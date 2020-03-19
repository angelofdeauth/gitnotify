/*
 * @File:     main_test.go
 * @Created:  2020-03-18 21:48:02
 * @Modified: 2020-03-18 21:48:06
 * @Author:   Antonio Escalera
 * @Commiter: Antonio Escalera
 * @Mail:     aj@angelofdeauth.host
 * @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>
 */

package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ret := m.Run()
	os.Exit(ret)
}
