// Copyright 2020 Jefferson Marchetti Ferreira. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"jeffersonmarchetti.com/fuegoquasar/server/routes"
	topsecrect_svc "jeffersonmarchetti.com/fuegoquasar/server/routes/topsecret"
	topsecrectsplit_svc "jeffersonmarchetti.com/fuegoquasar/server/routes/topsecretsplit"
)

const defaultPort = "8080"

// Start Iniciamos el servidor HTTP, escuchando en un puerto en particular
func Start() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Registramos las rutas que necesitamos
	r := mux.NewRouter()
	r.HandleFunc("/", routes.Healthcheck).Methods("GET")
	r.HandleFunc("/topsecret", topsecrect_svc.TopSecretService).Methods("POST")
	r.HandleFunc("/topsecret_split/{satellite_name}", topsecrectsplit_svc.TopSecretSplitService).Methods("POST")
	r.HandleFunc("/topsecret_split", topsecrectsplit_svc.TopSecretSplitGetService).Methods("GET")
	http.Handle("/", r)

	// Iniciamos el servidor
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
