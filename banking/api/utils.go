package api

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		panic(err)
	}
}
