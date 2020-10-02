// Copyright 2020 Jefferson Marchetti Ferreira. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package topsecret

import "testing"

func Execute_topsecret_logic_test(t *testing.T) {

	var topsecret_body2 Satellites_json
	topsecret_body2.Satellites = append(topsecret_body2.Satellites, Satellites{Name: "Kenobi", Distance: 111.2, Message: []string{"esta", "", "", ""}})
	topsecret_body2.Satellites = append(topsecret_body2.Satellites, Satellites{Name: "Sato", Distance: 211, Message: []string{"", "", "e", ""}})
	topsecret_body2.Satellites = append(topsecret_body2.Satellites, Satellites{Name: "Skywalker", Distance: 11, Message: []string{"", "", "", "mensagem"}})

	var _, _, _, validtest = execute_topsecret_logic(topsecret_body2)

	if !validtest {
		t.Error()
	}
}
