// Copyright 2020 Jefferson Marchetti Ferreira. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package core

import (
	"strings"

	"jeffersonmarchetti.com/fuegoquasar/models"
)

func checkInputData(distances []float32) bool {
	if distances == nil {
		return false
	}

	if len(distances) < 3 {
		return false
	}

	return true
}

func GetLocation(distances ...float32) (x, y float32) {

	if checkInputData(distances) != true {
		return -1, -1
	}

	var r1 float32 = distances[0]
	var r2 float32 = distances[1]
	var r3 float32 = distances[2]

	var a1 = models.Kenobi.Longitude
	var a2 = models.Skywalker.Longitude
	var a3 = models.Sato.Longitude

	var b1 = models.Kenobi.Latitude
	var b2 = models.Skywalker.Latitude
	var b3 = models.Sato.Latitude

	var (
		a1Sq float32 = a1 * a1
		a2Sq float32 = a2 * a2
		a3Sq float32 = a3 * a3
		b1Sq float32 = b1 * b1
		b2Sq float32 = b2 * b2
		b3Sq float32 = b3 * b3
		r1Sq float32 = r1 * r1
		r2Sq float32 = r2 * r2
		r3Sq float32 = r3 * r3
	)

	var numerator1 float32 = (a2-a1)*(a3Sq+b3Sq-r3Sq) + (a1-a3)*(a2Sq+b2Sq-r2Sq) + (a3-a2)*(a1Sq+b1Sq-r1Sq)
	var denominator1 float32 = 2 * (b3*(a2-a1) + b2*(a1-a3) + b1*(a3-a2))
	var resY = numerator1 / denominator1
	var numerator2 float32 = r2Sq - r1Sq + a1Sq - a2Sq + b1Sq - b2Sq - 2*(b1-b2)*resY
	var denominator2 float32 = 2 * (a1 - a2)
	var resX = numerator2 / denominator2
	return resX, resY
}

func GetMessage(messages ...[]string) (msg string) {

	var msg_res []string

	if !checkMessageParameters(messages) {
		return ""
	}

	var bigArray = getBiggerMessage(messages)

	for i := 0; i < len(bigArray); i++ {
		var (
			strtemp1 = ""
			strtemp2 = ""
			strtemp3 = ""
		)

		if i < len(messages[0]) && len(messages[0][i]) > 0 {
			strtemp1 = messages[0][i]
		}
		if i < len(messages[1]) && len(messages[1][i]) > 0 {
			strtemp2 = messages[1][i]
		}
		if i < len(messages[2]) && len(messages[2][i]) > 0 {
			strtemp3 = messages[2][i]
		}

		if len(strtemp1) > 0 {
			msg_res = append(msg_res, strtemp1)
		}
		if len(strtemp2) > 0 && strtemp2 != strtemp1 {
			msg_res = append(msg_res, strtemp2)
		}
		if len(strtemp3) > 0 && strtemp3 != strtemp2 && strtemp3 != strtemp1 {
			msg_res = append(msg_res, strtemp3)
		}
	}

	return strings.Join(msg_res, " ")
}

func checkMessageParameters(messages [][]string) bool {

	if messages == nil {
		return false
	} else if len(messages) < 3 {
		return false
	} else if messages[0] == nil && len(messages[0]) < 0 {
		return false
	} else if messages[1] == nil && len(messages[1]) < 0 {
		return false
	} else if messages[2] == nil && len(messages[2]) < 0 {
		return false
	}

	return true
}

func getBiggerMessage(messages [][]string) []string {

	if len(messages[0]) > len(messages[1]) && len(messages[0]) > len(messages[2]) {
		return messages[0]
	} else if len(messages[1]) > len(messages[0]) && len(messages[1]) > len(messages[2]) {
		return messages[1]
	} else if len(messages[2]) > len(messages[0]) && len(messages[2]) > len(messages[1]) {
		return messages[2]
	}

	if len(messages[0]) == len(messages[1]) && len(messages[0]) == len(messages[2]) {
		return messages[0]
	} else if len(messages[1]) == len(messages[2]) {
		return messages[1]
	} else {
		return messages[2]
	}
}
