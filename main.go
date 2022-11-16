package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"public-api/src"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", src.HomeHandler)
	srv := &http.Server{
		Handler: r,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
