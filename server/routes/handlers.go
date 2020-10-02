// Copyright 2020 Jefferson Marchetti Ferreira. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package routes

import (
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	json := simplejson.New()
	json.Set("Operación Fuego de Quasar", "Started")
	json.Set("message", "running...")

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(http.StatusOK)
}
