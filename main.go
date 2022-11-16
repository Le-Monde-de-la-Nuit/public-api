package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"public-api/src/database"
	"public-api/src/database/postgres"
	"public-api/src/handler"
	"strconv"
	"time"
)

func main() {
	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(os.Args)
		panic(err)
	}
	database.PublicCredentials = &database.Credentials{
		Host:         os.Args[1],
		Port:         port,
		User:         os.Args[3],
		Password:     os.Args[4],
		DatabaseName: os.Args[5],
	}
	check()
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler)
	r.HandleFunc("/user", handler.UserHandler)
	r.HandleFunc("/role", handler.RoleHandler)
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
