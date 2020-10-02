// Copyright 2020 Jefferson Marchetti Ferreira. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"jeffersonmarchetti.com/fuegoquasar/core"
)

func GetLocation_test(t *testing.T) {

	var long, lat = core.GetLocation(100.0, 115.5, 142.7)

	if long == -1 || lat == -1 {
		t.Error()
	}
}

func GetMessage_test(t *testing.T) {

	var msg = core.GetMessage([]string{"This", "", "", "message"}, []string{"", "a", "new", "message", "Z"}, []string{"This", "", "", "", "W"})

	if len(msg) <= 0 {
		t.Error()
	}
}
