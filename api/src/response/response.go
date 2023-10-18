package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if data != nil {
		log.Println(data)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			ERROR(w, statusCode, err)
		}
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		ERROR string `json:"error"`
	}{
		ERROR: err.Error(),
	})
}
