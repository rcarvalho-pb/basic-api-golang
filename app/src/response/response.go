package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			ERROR(w, http.StatusInternalServerError, err)
		}
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	JSON(w, statusCode, struct{
		ERROR string `json:"error"`
	}{
		ERROR: err.Error(),
	})
}