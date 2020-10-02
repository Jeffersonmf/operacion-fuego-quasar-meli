// Copyright 2020 Jefferson Marchetti Ferreira. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package main

import (
	"runtime"

	"jeffersonmarchetti.com/fuegoquasar/server"
)

func main() {

	//Setup the max counter of parallels Processors.
	runtime.GOMAXPROCS(1)

	server.Start()
}
