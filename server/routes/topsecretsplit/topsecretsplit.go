// Copyright 2020 Jefferson Marchetti Ferreira. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package topsecretsplit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"jeffersonmarchetti.com/fuegoquasar/core"
)

var Cache = cache.New(5*time.Minute, 5*time.Minute)

func setupHeader(res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(res).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	res.Header().Set("Content-Type", "application/json")
}

func SetCache(key string, emp interface{}) bool {
	Cache.Set(key, emp, cache.NoExpiration)
	return true
}

func GetCache(key string) (Satellites, bool) {
	var emp Satellites
	var found bool
	data, found := Cache.Get(key)
	if found {
		emp = data.(Satellites)
	}
	return emp, found
}

func TopSecretSplitGetService(res http.ResponseWriter, req *http.Request) {

	setupHeader(res)

	if req.Method != "GET" {
		log.Printf("HTTP Method not supported")
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, `{"message": "Method is not supported."}`)
		return
	}

	var kenobi, found1 = GetCache("kenobi")
	var skywalker, found2 = GetCache("skywalker")
	var sato, found3 = GetCache("sato")
	var (
		x       float32 = -1
		y       float32 = -1
		message string  = ""
	)

	if found1 && found2 && found3 {
		x, y, message = execute_topsecret_logic(kenobi, skywalker, sato)

		result := []byte(fmt.Sprintf("{\"position\": {\"x\": %.2f,\"y\": %.2f},\"message\": \"%s\"}", x, y, message))
		res.Write(result)
		res.WriteHeader(http.StatusOK)
	} else {
		result := []byte(`We have not yet collected enough data for this operation 
		... Satellites needed are: kenobi, skywalker and sato`)
		res.Write(result)
		res.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
	}
}

func TopSecretSplitService(res http.ResponseWriter, req *http.Request) {

	var topsecretBodySplit Satellites

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

	err := json.Unmarshal(body, &topsecretBodySplit)
	if err != nil {
		fmt.Println(err.Error())
	}

	var satelliteName = mux.Vars(req)["satellite_name"]
	println(satelliteName)
	topsecretBodySplit.Name = satelliteName

	switch strings.ToLower(satelliteName) {
	case "kenobi":
		SetCache("kenobi", topsecretBodySplit)
	case "skywalker":
		SetCache("skywalker", topsecretBodySplit)
	case "sato":
		SetCache("sato", topsecretBodySplit)
	default:
		res.Write([]byte(`The informed Satellite is not valid ... 
		Only satellites accepted in this version are: kenobi, skywalker and sato`))
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Write([]byte(`Satellite data ` + satelliteName + ` was received successfully:
	After informing of the remaining satellites, use the topsecret_split (GET) method to obtain the results.`))
	res.WriteHeader(http.StatusOK)
}

func execute_topsecret_logic(kenobi Satellites, skywalker Satellites, sato Satellites) (float32, float32, string) {

	var x, y = core.GetLocation(kenobi.Distance, skywalker.Distance, sato.Distance)
	var message = core.GetMessage(kenobi.Message, skywalker.Message, sato.Message)

	return x, y, message
}
