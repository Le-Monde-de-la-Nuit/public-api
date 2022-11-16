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
