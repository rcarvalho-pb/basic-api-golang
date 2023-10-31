package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
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

	url := fmt.Sprintf("%s/auth/login", config.ApiUrl)

	res, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		fmt.Println("err request")
		response.JSON(w, http.StatusInternalServerError, response.APIError{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		fmt.Println("err status code")
		response.StatusCodeErrorTreatment(w, res)
		return
	}

	var auth models.AuthDTO

	if err = json.NewDecoder(res.Body).Decode(&auth); err != nil {
		fmt.Println("err inside json decoder")
		response.JSON(w, http.StatusUnprocessableEntity, response.APIError{Error: err.Error()})
		return
	}
	
	if err = cookies.Save(w, fmt.Sprintf("%d",auth.ID), auth.Token); err != nil {
		fmt.Println("err inside save cookies")
		response.JSON(w, http.StatusUnprocessableEntity, response.APIError{Error: err.Error()})
		return
	}
	
	response.JSON(w, http.StatusOK, nil)
}
