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
	"time"
)

func main() {
	database.PublicCredentials = &database.Credentials{
		User:     os.Args[1],
		Password: os.Args[2],
	}
	check()
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler)
	r.HandleFunc("/member", handler.UserHandler)
	r.HandleFunc("/members", handler.UserHandler)
	r.HandleFunc("/role", handler.RoleHandler)
	r.HandleFunc("/roles", handler.RoleHandler)
	srv := &http.Server{
		Handler: r,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("API is running")
	log.Fatal(srv.ListenAndServe())
}

func check() {
	credentials := database.PublicCredentials
	// Roles before user
	db, err := credentials.Connect("postgres", "members")
	if err != nil {
		panic(err)
	}
	pg := postgres.Postgres{Db: db}
	pg.GenerateRolesTable().GenerateUsersTable()
	// Places and types before actions
	db, err = credentials.Connect("postgres", "actions")
	if err != nil {
		panic(err)
	}
	pg = postgres.Postgres{Db: db}
	pg.GeneratePlacesTable().GenerateTypesTable().GenerateActionsTable()
	// Posts and Tags before PostTags
	db, err = credentials.Connect("postgres", "posts")
	if err != nil {
		panic(err)
	}
	pg = postgres.Postgres{Db: db}
	pg.GeneratePostsTable().GenerateTagsTable().GeneratePostTagsTable()
	defer db.Close()
}
