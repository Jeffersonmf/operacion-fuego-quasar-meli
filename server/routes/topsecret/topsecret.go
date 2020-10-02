// Copyright 2020 Jefferson Marchetti Ferreira. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package topsecret

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"jeffersonmarchetti.com/fuegoquasar/core"
)

func setupHeader(res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(res).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	res.Header().Set("Content-Type", "application/json")
}

func TopSecretService(res http.ResponseWriter, req *http.Request) {

	var topsecret_body2 Satellites_json

	setupHeader(res)

	if req.Method != "POST" {
		log.Printf("HTTP Method not supported")
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, `{"message": "Method is not supported."}`)
		return
	}

	body, err4 := ioutil.ReadAll(req.Body)
	if err4 != nil {
		log.Printf("Error found trying to read the body")

		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.Unmarshal(body, &topsecret_body2)
	if err != nil {
		fmt.Println(err.Error())
	}

	var x, y, message, validResult = execute_topsecret_logic(topsecret_body2)

	if validResult {
		result := []byte(fmt.Sprintf("{\"position\": {\"x\": %.2f,\"y\": %.2f},\"message\": \"%s\"}", x, y, message))
		res.Write(result)
		res.WriteHeader(http.StatusOK)
	} else {
		result := []byte(`The informed Satellite is not valid ... 
		Only satellites accepted in this version are: kenobi, skywalker and sato`)
		res.WriteHeader(http.StatusBadRequest)
		res.Write(result)
	}
}

func execute_topsecret_logic(topsecret_body Satellites_json) (float32, float32, string, bool) {

	if topsecret_body.Satellites == nil || len(topsecret_body.Satellites) < 3 {
		return -1, -1, "", false
	}

	var (
		r1      float32 = 0
		r2      float32 = 0
		r3      float32 = 0
		x       float32 = 0
		y       float32 = 0
		message string  = ""
	)

	messages := make(map[string][]string)
	itensCollected := 0

	for _, item := range topsecret_body.Satellites {
		switch strings.ToLower(item.Name) {
		case "kenobi":
			r1 = item.Distance
			messages["kenobi"] = item.Message
			itensCollected++
		case "skywalker":
			r2 = item.Distance
			messages["skywalker"] = item.Message
			itensCollected++
		case "sato":
			r3 = item.Distance
			messages["sato"] = item.Message
			itensCollected++
		}
	}

	if itensCollected < 3 {
		return -1, -1, message, false
	}

	x, y = core.GetLocation(r1, r2, r3)
	message = core.GetMessage(messages["kenobi"], messages["skywalker"], messages["sato"])

	return x, y, message, true
}
