// Copyright 2020 Jefferson Marchetti Ferreira. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package topsecretsplit

import "testing"

func execute_topsecretsplit_logic_test(t *testing.T) {

	//execute_topsecret_logic()
}

func Execute_topsecret_logic_test(t *testing.T) {

	var kenobiBody Satellites
	var skywalkerBody Satellites
	var satoBody Satellites

	kenobiBody = Satellites{Name: "Kenobi", Distance: 111.2, Message: []string{"esta", "", "", ""}}
	skywalkerBody = Satellites{Name: "Skywalker", Distance: 111.2, Message: []string{"esta", "", "", ""}}
	satoBody = Satellites{Name: "Sato", Distance: 111.2, Message: []string{"esta", "", "", ""}}

	var x, y, messages = execute_topsecret_logic(kenobiBody, skywalkerBody, satoBody)

	if x == -1 || y == -1 || messages == "" {
		t.Error()
	}
}
