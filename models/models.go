// Copyright 2020 Jefferson Marchetti Ferreira. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package models

type Satelite struct {
	Longitude float32
	Latitude  float32
}

var (
	Kenobi    = Satelite{kenobiX, kenobiY}
	Skywalker = Satelite{skywalkerX, skywalkerY}
	Sato      = Satelite{satoX, satoY}
)

const kenobiX = -500
const kenobiY = -200

const skywalkerX = 100
const skywalkerY = -100

const satoX = 500
const satoY = 100
