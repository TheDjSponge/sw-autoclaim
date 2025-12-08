package main

import (
	"encoding/json"
	"log"
	"net/http"
)


type JsonResponse struct{
		Status int `json:"status"`
		Message string `json:"message"` 
	}

func respondWithError(w http.ResponseWriter, code int, msg string) {

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

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	dat, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, 500, "Couldn't marshal response body")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}