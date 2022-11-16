package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"public-api/src/database"
	"public-api/src/database/postgres"
	"public-api/src/handler"
	"time"
)

func main() {
	database.PublicCredentials = database.ParseConnectionString(os.Args[1])
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler)
	srv := &http.Server{
		Handler: r,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func check() {
	db, err := database.Connect("postgres", database.PublicCredentials)
	if err != nil {
		panic(err)
	}
	// Roles before user
	postgres.GenerateRolesTable(db)
	postgres.GenerateUsersTable(db)
	// Places and types before actions
	postgres.GeneratePlacesTable(db)
	postgres.GenerateTypesTable(db)
	postgres.GenerateActionsTable(db)
	defer db.Close()
}
