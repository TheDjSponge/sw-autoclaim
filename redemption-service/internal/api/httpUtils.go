package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func RespondWithMessage(w http.ResponseWriter, code int, msg string) {

	retError := JsonResponse{Status: code, Message: msg}
	byteError, err := json.Marshal(retError)
	if err != nil {
		log.Printf("Error marshalling response JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(byteError)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload any) {

	dat, err := json.Marshal(payload)
	if err != nil {
		RespondWithMessage(w, 500, "Couldn't marshal response body")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}


func GetServerHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
