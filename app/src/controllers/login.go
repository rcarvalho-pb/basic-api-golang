package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}