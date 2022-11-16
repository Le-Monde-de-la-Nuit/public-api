package handler

import (
	"encoding/json"
	"net/http"
	"public-api/src/response"
)

func defaultHandler(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
}

func writeResponse(w http.ResponseWriter, resp *response.Response) error {
	w.WriteHeader(resp.StatusCode)
	return json.NewEncoder(w).Encode(resp)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	defaultHandler(r)
	resp := response.Response{StatusCode: 200, Message: "Welcome to the API"}
	err := writeResponse(w, &resp)
	if err != nil {
		panic(err)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	handler(w, r, nil, postUserHandler)
}

func RoleHandler(w http.ResponseWriter, r *http.Request) {
	handler(w, r, nil, postRoleHandler)
}

func handler(w http.ResponseWriter, r *http.Request, g func(http.ResponseWriter, *http.Request, *response.Response),
	p func(http.ResponseWriter, *http.Request, *response.Response)) {
	defaultHandler(r)
	resp := response.Response{StatusCode: 200}
	switch r.Method {
	case http.MethodGet:
		resp.Message = "Get role" // use g later
	case http.MethodPost:
		p(w, r, &resp)
	}
	err := writeResponse(w, &resp)
	if err != nil {
		panic(err)
	}
}
