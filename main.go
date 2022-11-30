package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"public-api/src/database"
	"public-api/src/database/mariadb"
	"public-api/src/handler"
	"time"
)

func main() {
	database.PublicCredentials = &database.Credentials{
		User:     os.Args[1],
		Password: os.Args[2],
	}
	log.Println("sleeping for 10 seconds")
	time.Sleep(10 * time.Second)
	log.Println("check the databases")
	check()
	log.Println("check finished")

	r := mux.NewRouter()

	handle(r)
	log.Println("starting server")
	srv := &http.Server{
		Handler: r,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("API is running")
	log.Fatal(srv.ListenAndServe())
}

func handle(r *mux.Router) {
	// Members
	r.HandleFunc("/", handler.HomeHandler)
	r.HandleFunc("/member", handler.UserHandler)
	r.HandleFunc("/members", handler.UserHandler)
	r.HandleFunc("/role", handler.RoleHandler)
	r.HandleFunc("/roles", handler.RoleHandler)

	// Actions
	// TODO

	// Posts
	r.HandleFunc("/post", handler.PostHandler)
	r.HandleFunc("/posts", handler.PostHandler)
	r.HandleFunc("/tag", handler.TagHandler)
	r.HandleFunc("/tags", handler.TagHandler)
}

func check() {
	credentials := database.PublicCredentials
	// Roles before user
	db, err := credentials.Connect("mysql", "members")
	if err != nil {
		panic(err)
	}
	pg := mariadb.Mariadb{Db: db}
	pg.GenerateRolesTable().GenerateUsersTable()
	// Places and types before actions
	db, err = credentials.Connect("mysql", "actions")
	if err != nil {
		panic(err)
	}
	pg = mariadb.Mariadb{Db: db}
	pg.GeneratePlacesTable().GenerateTypesTable().GenerateActionsTable()
	// Posts and Tags before PostTags
	db, err = credentials.Connect("mysql", "posts")
	if err != nil {
		panic(err)
	}
	pg = mariadb.Mariadb{Db: db}
	pg.GeneratePostsTable().GenerateTagsTable().GeneratePostTagsTable()
	defer db.Close()
}
