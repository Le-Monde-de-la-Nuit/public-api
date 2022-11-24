package handler

import (
	"log"
	"net/http"
	"public-api/src/database"
	"public-api/src/database/postgres"
	"public-api/src/response"
)

func postUserHandler(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	data := postgres.User{}
	err := ParseBody(r, data)
	if err != nil {
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "Bad request"
		w.WriteHeader(resp.StatusCode)
		log.Fatal(err)
		return
	}
	postHandler(w, resp, "INSERT INTO users (id, role_id, discord) VALUES ($1, $2, $3)", data.ID, data.RoleID,
		data.Discord)
}

func postRoleHandler(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	data := postgres.Role{}
	err := ParseBody(r, &data)
	if err != nil {
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "Bad request"
		w.WriteHeader(resp.StatusCode)
		log.Fatal(err)
		return
	}
	postHandler(w, resp, "INSERT INTO roles (name, description) VALUES ($1, $2)", data.Name, data.Description)
}

func postHandler(w http.ResponseWriter, resp *response.Response, q string, values ...any) {
	db, err := database.Connect("postgres", database.PublicCredentials)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Message = "Internal server error"
		w.WriteHeader(resp.StatusCode)
		log.Fatal(err)
		return
	}
	defer db.Close()
	postgres.Insert(db, q, values...)

	resp.StatusCode = http.StatusCreated
	resp.Message = "Created"
	resp.Data = values
	w.WriteHeader(resp.StatusCode)
}
