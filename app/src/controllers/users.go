package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/response"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"Name":     r.FormValue("Name"),
		"Email":    r.FormValue("Email"),
		"Nick":     r.FormValue("Nick"),
		"Password": r.FormValue("Password"),
	})
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
	}

	res, err := http.Post("http://localhost:5000/users", "application/json", bytes.NewBuffer(user))
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
	}
	defer res.Body.Close()

	response.JSON(w, res.StatusCode, nil)
}
