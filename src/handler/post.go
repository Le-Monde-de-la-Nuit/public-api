package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"public-api/src/database"
	"public-api/src/database/postgres"
	"public-api/src/response"
)

func postUserHandler(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	data := postgres.User{}
	postHandler(w, r, resp, postgres.User{}, "INSERT INTO users (id, role_id, discord) VALUES ($1, $2, $3)", data.ID,
		data.Role.ID, data.Discord)
}

func postRoleHandler(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	data := postgres.Role{}
	postHandler(w, r, resp, postgres.Role{}, "INSERT INTO roles (name, description) VALUES ($1, $2)", data.Name,
		data.Description)
}

func postHandler(w http.ResponseWriter, r *http.Request, resp *response.Response, data interface{}, q string, values ...interface{}) {
	db, err := database.Connect("postgres", database.PublicCredentials)
	if err != nil {
		resp.StatusCode = 500
		resp.Message = "Internal server error"
		w.WriteHeader(resp.StatusCode)
		log.Fatal(err)
		return
	}
	defer db.Close()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		resp.StatusCode = 500
		resp.Message = "Internal server error"
		w.WriteHeader(resp.StatusCode)
		log.Fatal(err)
		return
	}
	err = json.Unmarshal(content, &data)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = "Bad request"
		w.WriteHeader(resp.StatusCode)
		log.Fatal(err)
		return
	}

	result := postgres.Insert(db, q, values...)
	lastId, _ := result.LastInsertId()
	defData := struct {
		ID      int64       `json:"id"`
		Content interface{} `json:"content"`
	}{
		ID:      lastId,
		Content: data,
	}

	resp.StatusCode = 201
	resp.Message = "Created"
	resp.Data = defData
	w.WriteHeader(resp.StatusCode)
}
