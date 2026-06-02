package main

import (
	"encoding/json"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, msg string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	resp := errorResponse{
		Error: msg,
	}
	dat, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)

}
func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parametres struct {
		Body string `json:"body"`
	}
	params := parametres{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 500, "Something went wrong")
		return
	}
	if len(params.Body) > 140 {
		responseWithError(w, 400, "Chrips too long")
		return
	}

	type response struct {
		Valid bool `json:"valid`
	}
	responseWithJSON(w, 200, response{
		Valid: true,
	})
}
