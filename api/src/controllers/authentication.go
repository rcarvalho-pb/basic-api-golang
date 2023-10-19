package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	token, err := authentication.CreateToken(uint32(user.ID))
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusAccepted, struct {
		Token  string `json:"token"`
		UserId uint32 `json:"userId"`
	}{
		Token:  token,
		UserId: uint32(user.ID),
	})
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	id, err := authentication.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	followedId, err := strconv.ParseUint(params["userId"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	db, queries, err := db.Conn()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(queries)

	if err = repository.FollowUser(id, uint32(followedId)); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}
