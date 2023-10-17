package controllers

import (
	"api/src/db"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

type Auth struct {
	EmailOrNick string `json:"emailOrNick"`
	Password    string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var auth Auth
	if err = json.Unmarshal(requestBody, &auth); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	sql, queries, err := db.Conn()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer sql.Close()

	repository := repositories.NewUserRepository(queries)

	user, err := repository.GetUserByEmailOrNick(auth.EmailOrNick)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err = security.ValidatePassword(auth.Password, user.Password); err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	response.JSON(w, http.StatusAccepted, struct{
		Login string
	}{
		Login: "Logged sucessifully",
	})
}
