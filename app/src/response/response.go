package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIError struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func StatusCodeErrorTreatment(w http.ResponseWriter, r *http.Response) {
	var err APIError

	if err := json.NewDecoder(r.Body).Decode(&err); err != nil {
		log.Fatal(err)
	}
	JSON(w, r.StatusCode, err)
}