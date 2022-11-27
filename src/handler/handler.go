package handler

import (
	"encoding/json"
	"net/http"
	"public-api/src/response"
	"strings"
)

type handle struct {
	w http.ResponseWriter
	r *http.Request
	g func(http.ResponseWriter, *http.Request, *response.Response)
	p func(http.ResponseWriter, *http.Request, *response.Response)
}

func defaultHandler(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
}

func writeResponse(w http.ResponseWriter, resp *response.Response) error {
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
	}
	return json.NewEncoder(w).Encode(&resp)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	defaultHandler(r)
	resp := response.Response{StatusCode: http.StatusOK, Message: "Welcome to the API"}
	err := writeResponse(w, &resp)
	if err != nil {
		panic(err)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	handle := handle{w: w, r: r, g: getUserHandler, p: postUserHandler}
	handle.handler()
}

func RoleHandler(w http.ResponseWriter, r *http.Request) {
	handle := handle{w: w, r: r, g: getRoleHandler, p: postRoleHandler}
	handle.handler()
}

func (h *handle) handler() {
	r := h.r
	w := h.w
	g := h.g
	p := h.p
	defaultHandler(r)
	resp := response.Response{StatusCode: http.StatusOK}
	switch r.Method {
	case http.MethodGet:
		g(w, r, &resp)
	case http.MethodPost:
		p(w, r, &resp)
	}
	err := writeResponse(w, &resp)
	if err != nil {
		panic(err)
	}
}

func ParseQueryInURI(uri string) map[string]string {
	query := strings.Split(uri, "?")
	if len(query) < 2 {
		return nil
	}
	queries := strings.Split(query[1], "&")
	data := map[string]string{}

	for _, v := range queries {
		splited := strings.Split(v, "=")
		if len(splited) < 2 {
			continue
		}
		data[splited[0]] = splited[1]
	}
	if len(data) == 0 {
		return nil
	}
	return data
}

func ParseBody(r *http.Request, i interface{}) error {
	err := json.NewDecoder(r.Body).Decode(i)
	if err != nil {
		return err
	}
	return nil
}
