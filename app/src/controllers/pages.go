package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/request"
	"webapp/src/response"
	"webapp/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, _ *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoadUserRegisterPage(w http.ResponseWriter, _ *http.Request) {
	utils.ExecuteTemplate(w, "register.html", nil)
}

func LoadHome(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publications", config.ApiUrl)

	res, err := request.Request(r, http.MethodGet, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.APIError{Error: err.Error()})
		return
	}

	if res.StatusCode >= 400 {
		response.StatusCodeErrorTreatment(w, res)
		return
	}

	var publications []models.Publication
	if err = json.NewDecoder(res.Body).Decode(&publications); err != nil {
		fmt.Println(publications)
		response.JSON(w, http.StatusUnprocessableEntity, response.APIError{Error: err.Error()})
		return
	}
	utils.ExecuteTemplate(w, "home.html", nil)
}
