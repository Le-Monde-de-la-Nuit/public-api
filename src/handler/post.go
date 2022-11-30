package handler

import (
	"database/sql"
	"log"
	"net/http"
	"public-api/src/database"
	"public-api/src/database/mariadb"
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
	data := mariadb.User{}
	err := ParseBody(r, &data)
	if err != nil {
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "Bad request"
		log.Println(err)
		return
	}
	h.postHandler("INSERT INTO users (id, role_id, discord) VALUES (?, ?, ?)", data.ID, data.RoleID,
		data.Discord)
	resp.Data = data
}

func postRoleHandler(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	h := postHandle{w: w, r: r, resp: resp, dbName: "members"}
	data := mariadb.Role{}
	err := ParseBody(r, &data)
	if err != nil {
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "Bad request"
		log.Println(err)
		return
	}
	result := h.postHandler("INSERT INTO roles (name, description) VALUES (?, ?)", data.Name, data.Description)
	data.ID, err = result.LastInsertId()
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Message = "Internal server error"
		log.Println(err)
		return
	}
	resp.Data = data
}

func postPostHandler(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	h := postHandle{w: w, r: r, resp: resp, dbName: "posts"}
	data := mariadb.NewPost{}
	err := ParseBody(r, &data)
	if err != nil {
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "Bad request"
		log.Println(err)
		return
	}
	result := h.postHandler("INSERT INTO posts (title, content) VALUES (?, ?)", data.Title, data.Content)
	if h.resp.StatusCode != http.StatusCreated {
		return
	}
	db, err := h.getDatabase()
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Message = "Internal Server Error while connecting to the database"
		log.Println(err)
		return
	}
	defer db.Close()
	id, err := result.LastInsertId()
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Message = "Internal server error"
		log.Println(err)
		return
	}
	data.ID = id
	for _, v := range data.Tag {
		h.postHandler("INSERT INTO post_tags (post_id, tag_id) VALUES (?, ?)", id, v)
		if h.resp.StatusCode != http.StatusCreated {
			return
		}
	}
	finalData := data.ToPost()
	t, err := database.GetTagsForPost(db, data.ID)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Message = "Internal server error"
		log.Println(err)
		return
	}
	finalData.Tag = t
	h.resp.Data = finalData
}

func postTagHandler(w http.ResponseWriter, r *http.Request, resp *response.Response) {
	h := postHandle{w: w, r: r, resp: resp, dbName: "posts"}
	data := mariadb.Tag{}
	err := ParseBody(r, &data)
	if err != nil {
		resp.StatusCode = http.StatusBadRequest
		resp.Message = "Bad request"
		log.Println(err)
		return
	}
	result := h.postHandler("INSERT INTO tags (name) VALUES (?)", data.Name)
	data.ID, err = result.LastInsertId()
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Message = "Internal server error"
		log.Println(err)
		return
	}
	resp.Data = data
}

func (h *postHandle) postHandler(q string, values ...any) sql.Result {
	resp := h.resp
	db, err := h.getDatabase()
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Message = "Internal Server Error while connecting to the database"
		log.Println(err)
		return nil
	}
	defer db.Close()
	pg := mariadb.Mariadb{Db: db}
	result := pg.Insert(q, values...)

	resp.StatusCode = http.StatusCreated
	resp.Message = "Created"
	return result
}

func (h *postHandle) getDatabase() (*sql.DB, error) {
	c := database.PublicCredentials
	db, err := c.Connect("mysql", h.dbName)
	if err != nil {
		return nil, err
	}
	return db, nil
}
