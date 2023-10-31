package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/request"
	"webapp/src/response"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publication, err := json.Marshal(map[string]string{
		"Title": r.FormValue("Title"),
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