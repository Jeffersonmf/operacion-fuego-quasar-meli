// Copyright 2020 Jefferson Marchetti Ferreira. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package utils

func IsIntegral(val float32) bool {
	return val == float32(int(val))
}
