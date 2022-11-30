package handler

import (
	"database/sql"
	"net/http"
	"public-api/src/database"
	"public-api/src/database/mariadb"
	"public-api/src/response"
)

type getHandle struct {
	w      http.ResponseWriter
	r      *http.Request
	resp   *response.Response
	pg     mariadb.Mariadb
	dbName string
}

func getUserHandler(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	h := getHandle{w: w, r: r, resp: resp}
	h.initializeDbForGet("members")
	result, err := h.getHandler("users")
	defer h.pg.Db.Close()

	if h.manageErrors(result, err) {
		return
	}

	var users []mariadb.User
	for result.Next() {
		var user mariadb.User
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
	h.initializeDbForGet("members")
	result, err := h.getHandler("roles")
	defer h.pg.Db.Close()

	if h.manageErrors(result, err) {
		return
	}

	var roles []mariadb.Role
	for result.Next() {
		var role mariadb.Role
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

func getPostHandler(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	h := getHandle{w: w, r: r, resp: resp}
	h.initializeDbForGet("posts")
	result, err := h.getHandler("posts")
	defer h.pg.Db.Close()

	if h.manageErrors(result, err) {
		return
	}

	var rawPosts []mariadb.RawPost
	for result.Next() {
		var post mariadb.RawPost
		err := result.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			panic(err)
		}
		rawPosts = append(rawPosts, post)
	}
	resp.StatusCode = http.StatusOK
	resp.Message = "OK"
	data := make([]mariadb.Post, len(rawPosts))
	for i := range rawPosts {
		data = append(data, rawPosts[i].ToPost())
	}
	for i := range data {
		t, err := database.GetTagsForPost(h.pg.Db, data[i].ID)
		if err != nil {
			panic(err)
		}
		data[i].Tag = t
	}
	resp.Data = data
}

func getTagHandler(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	h := getHandle{w: w, r: r, resp: resp, dbName: "posts"}
	h.initializeDbForGet("posts")
	result, err := h.getHandler("tags")
	defer h.pg.Db.Close()

	if h.manageErrors(result, err) {
		return
	}

	var tags []mariadb.Tag
	for result.Next() {
		var tag mariadb.Tag
		err := result.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err)
		}
		tags = append(tags, tag)
	}
	resp.StatusCode = http.StatusOK
	resp.Message = "OK"
	resp.Data = tags
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
	return h.pg.Get("SELECT * FROM "+table+" WHERE id = ?", id)
}

func (h *getHandle) getAllHandler(table string) (*sql.Rows, error) {
	return h.pg.Get("SELECT * FROM " + table)
}

func (h *getHandle) initializeDbForGet(dbName string) {
	c := database.PublicCredentials
	db, err := c.Connect("mysql", dbName)
	if err != nil {
		panic(err)
	}
	h.pg = mariadb.Mariadb{Db: db}
	h.dbName = dbName
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
