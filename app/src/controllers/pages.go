package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Login Page")
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoadUserRegisterPage(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Register Page")
	utils.ExecuteTemplate(w, "register.html", nil)
}

func LoadHome(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Home Page")
	// url := fmt.Sprintf("%s/publications", config.ApiUrl)
	// fmt.Println(url)
	// res, err := request.Request(r, http.MethodGet, url, nil)

	// fmt.Println(res.StatusCode, err)

	utils.ExecuteTemplate(w, "home.html", nil)
}
