package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
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
		log.Fatal(err)
	}

	res, err := http.Post(fmt.Sprintf("%s/users", config.ApiUrl), "application/json", bytes.NewBuffer(user))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.APIError{Error: err.Error()})
		return
	}
	defer res.Body.Close()
	
	if res.StatusCode >= 400 {
		response.StatusCodeErrorTreatment(w, res)
		return
	}

	response.JSON(w, res.StatusCode, nil)
}
