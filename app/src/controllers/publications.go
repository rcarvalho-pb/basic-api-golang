package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/request"
	"webapp/src/response"

	"github.com/gorilla/mux"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publication, err := json.Marshal(map[string]string{
		"Title":   r.FormValue("Title"),
		"Content": r.FormValue("Content"),
	})
	if err != nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf("%s/publications", config.ApiUrl)

	res, err := request.Request(r, http.MethodPost, url, bytes.NewBuffer(publication))
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIError{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.StatusCodeErrorTreatment(w, res)
		return
	}

	response.JSON(w, res.StatusCode, nil)
}

func LikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationId, err := strconv.ParseInt(params["publicationId"], 10, 32)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d/like", config.ApiUrl, publicationId)

	res, err := request.Request(r, http.MethodPatch, url, nil)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIError{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.StatusCodeErrorTreatment(w, res)
		return
	}

	response.JSON(w, res.StatusCode, nil)
}

func DislikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationId, err := strconv.ParseInt(params["publicationId"], 10, 32)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d/dislike", config.ApiUrl, publicationId)

	res, err := request.Request(r, http.MethodPatch, url, nil)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIError{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.StatusCodeErrorTreatment(w, res)
		return
	}

	response.JSON(w, res.StatusCode, nil)
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationId, err := strconv.ParseInt(params["publicationId"], 10, 32)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, err)
		return
	}

	r.ParseForm()

	publication, err := json.Marshal(map[string]string{
		"Title": r.FormValue("title"),
		"Content": r.FormValue("content"),
	})
	if err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.ApiUrl, publicationId)
	res, err := request.Request(r, http.MethodPut, url, bytes.NewBuffer(publication))
	if err != nil {
		fmt.Println(err)
		response.JSON(w, http.StatusBadRequest, response.APIError{Error: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.StatusCodeErrorTreatment(w, res)
		return
	}

	response.JSON(w, res.StatusCode, nil)
}

func DeletetePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	publicationId, err := strconv.ParseInt(params["publicationId"], 10, 32)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, err)
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.ApiUrl, publicationId)

	res, err := request.Request(r, http.MethodDelete, url, nil)
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIError{ Error: err.Error() })
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		response.StatusCodeErrorTreatment(w, res)
		return
	}

	response.JSON(w, res.StatusCode, nil)
}