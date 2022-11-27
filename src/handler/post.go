package handler

import (
	"log"
	"net/http"
	"public-api/src/database"
	"public-api/src/database/postgres"
	"public-api/src/response"
)

type postHandle struct {
	w      http.ResponseWriter
	r      *http.Request
	resp   *response.Response
	dbName string
}

func postUserHandler(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	h := postHandle{w: w, r: r, resp: resp, dbName: "members"}
	data := postgres.User{}
	err := ParseBody(r, &data)
	if err != nil {
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "Bad request"
		w.WriteHeader(resp.StatusCode)
		log.Fatal(err)
		return
	}
	h.postHandler("INSERT INTO users (id, role_id, discord) VALUES ($1, $2, $3)", data.ID, data.RoleID,
		data.Discord)
}

func postRoleHandler(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	h := postHandle{w: w, r: r, resp: resp, dbName: "members"}
	data := postgres.Role{}
	err := ParseBody(r, &data)
	if err != nil {
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "Bad request"
		w.WriteHeader(resp.StatusCode)
		log.Fatal(err)
		return
	}
	h.postHandler("INSERT INTO roles (name, description) VALUES ($1, $2)", data.Name, data.Description)
}

func (h *postHandle) postHandler(q string, values ...any) {
	w := h.w
	resp := h.resp
	c := database.PublicCredentials
	db, err := c.Connect("postgres", c.GenerateConnectionString(h.dbName))
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Message = "Internal server error"
		w.WriteHeader(resp.StatusCode)
		log.Fatal(err)
		return
	}
	defer db.Close()
	pg := postgres.Postgres{Db: db}
	pg.Insert(q, values...)

	resp.StatusCode = http.StatusCreated
	resp.Message = "Created"
	resp.Data = values
	w.WriteHeader(resp.StatusCode)
}
