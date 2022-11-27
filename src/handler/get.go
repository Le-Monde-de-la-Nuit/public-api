package handler

import (
	"database/sql"
	"net/http"
	"public-api/src/database"
	"public-api/src/database/postgres"
	"public-api/src/response"
)

type getHandle struct {
	w    http.ResponseWriter
	r    *http.Request
	resp *response.Response
	pg   postgres.Postgres
}

func getUserHandler(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	h := getHandle{w: w, r: r, resp: resp}
	h.initializeDbForGet()
	result, err := h.getHandler("users")
	defer h.pg.Db.Close()

	if h.manageErrors(result, err) {
		return
	}

	var users []postgres.User
	for result.Next() {
		var user postgres.User
		err := result.Scan(&user.ID, &user.Discord, &user.RoleID)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	resp.StatusCode = http.StatusOK
	resp.Message = "OK"
	resp.Data = users
}

func getRoleHandler(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	h := getHandle{w: w, r: r, resp: resp}
	h.initializeDbForGet()
	result, err := h.getHandler("roles")
	defer h.pg.Db.Close()

	if h.manageErrors(result, err) {
		return
	}

	var roles []postgres.Role
	for result.Next() {
		var role postgres.Role
		err := result.Scan(&role.ID, &role.Name, &role.Description)
		if err != nil {
			panic(err)
		}
		roles = append(roles, role)
	}
	resp.StatusCode = http.StatusOK
	resp.Message = "OK"
	resp.Data = roles
}

func (h *getHandle) getHandler(table string) (*sql.Rows, error) {
	if h.r.RequestURI == "/"+table {
		return h.getAllHandler(table)
	}
	return h.getOneHandler(table)
}

func (h *getHandle) getOneHandler(table string) (*sql.Rows, error) {
	w := h.w
	r := h.r
	resp := h.resp
	queries := ParseQueryInURI(r.RequestURI)
	if queries == nil {
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "The query in the url is not valid"
		return nil, nil
	}
	id, ok := queries["id"]
	if !ok {
		resp.StatusCode = http.StatusBadRequest
		last := len(table) - 1
		resp.Message = "There is no id, if you don't want to take one " + table[0:last] + ", use /" + table + " instead"
		w.WriteHeader(resp.StatusCode)
		return nil, nil
	}
	return h.pg.Get("SELECT * FROM "+table+" WHERE id = $1", id)
}

func (h *getHandle) getAllHandler(table string) (*sql.Rows, error) {
	db, err := database.Connect("postgres", database.PublicCredentials)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return h.pg.Get("SELECT * FROM " + table)
}

func (h *getHandle) initializeDbForGet() {
	db, err := database.Connect("postgres", database.PublicCredentials)
	if err != nil {
		panic(err)
	}
	h.pg = postgres.Postgres{Db: db}
}

func (h *getHandle) manageErrors(result *sql.Rows, err error) bool {
	resp := h.resp
	if result == nil {
		if resp.Message != "" {
			return true
		}
		println("result is nil")
		resp.Data = nil
		resp.StatusCode = http.StatusOK
		resp.Message = "No result find"
		return true
	}
	if err != nil {
		println(err.Error())
		resp.Data = nil
		resp.StatusCode = http.StatusInternalServerError
		resp.Message = "Internal server error"
		return true
	}
	return false
}
