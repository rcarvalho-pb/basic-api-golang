package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/response"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"emailOrNick": r.FormValue("email"),
		"password":    r.FormValue("password"),
	})
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIError{Error: err.Error()})
		return
	}

	res, err := http.Post("http://localhost:3000/auth/login", "application/json", bytes.NewBuffer(user))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.APIError{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.StatusCodeErrorTreatment(w, res)
		return
	}

	var auth models.AuthDTO

	if err = json.NewDecoder(res.Body).Decode(&auth); err != nil {
		log.Println("Aqui")
		log.Println(auth)
		response.JSON(w, http.StatusUnprocessableEntity, response.APIError{Error: err.Error()})
		return
	}

	if err = cookies.Save(w, auth.ID, auth.Token); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.APIError{Error: err.Error()})
		return
	}

	response.JSON(w, http.StatusOK, nil)
}
