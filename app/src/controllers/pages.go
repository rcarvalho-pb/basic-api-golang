package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/request"
	"webapp/src/response"
	"webapp/src/utils"

	"github.com/gorilla/mux"
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
		response.JSON(w, http.StatusUnprocessableEntity, response.APIError{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseInt(cookie["id"], 10, 32)

	utils.ExecuteTemplate(w, "home.html", struct {
		Publications []models.Publication
		UserID       int32
	}{
		Publications: publications,
		UserID:       int32(userId),
	})
}

func LoadUpdatePublicationPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationId, err := strconv.ParseInt(params["publicationId"], 10, 32)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.ApiUrl, publicationId)

	res, err := request.Request(r, http.MethodGet, url, nil)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.StatusCodeErrorTreatment(w, res)
		return
	}

	var publication models.Publication

	if err = json.NewDecoder(res.Body).Decode(&publication); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.ExecuteTemplate(w, "update-publication.html", publication)
}
