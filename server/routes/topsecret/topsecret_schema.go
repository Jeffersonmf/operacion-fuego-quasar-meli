// Copyright 2020 Jefferson Marchetti Ferreira. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package topsecret

// Messages Mensajes recibidos del emisor en cada satelite
type Satellites struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type Satellites_json struct {
	Satellites []Satellites `json:"satellites"`
}
