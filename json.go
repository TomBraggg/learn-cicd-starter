package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	// Marshal the payload to JSON
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := `{"error":"Internal server error"}`
		if _, writeErr := w.Write([]byte(errorResponse)); writeErr != nil {
			log.Printf("Error writing error response: %s", writeErr)
		}
		return
	}

	// Write the status code
	w.WriteHeader(code)

	// Write the JSON response
	if _, writeErr := w.Write(dat); writeErr != nil {
		log.Printf("Error writing JSON response: %s", writeErr)
	}
}
