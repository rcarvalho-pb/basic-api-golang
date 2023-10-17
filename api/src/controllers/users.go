package controllers

import (
	"api/db/database"
	"api/src/db"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	var user database.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = user.PrepareUser(database.NewUser); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	sql, queries, err := db.Conn()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer sql.Close()

	repository := repositories.NewUserRepository(queries)

	userIndex, err := repository.CreateUser(user)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	user.ID = int32(userIndex)

	response.JSON(w, http.StatusOK, user)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	userIdentifier := strings.ToLower(r.URL.Query().Get("user"))
	sql, queries, err := db.Conn()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer sql.Close()

	repository := repositories.NewUserRepository(queries)

	users, err := repository.FindUser(userIdentifier)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["userId"])
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	sql, queries, err := db.Conn()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer sql.Close()

	repository := repositories.NewUserRepository(queries)

	user, err := repository.GetUserById(int32(userId))
	if err != nil {
		response.ERROR(w, http.StatusNotFound, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	var user database.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = user.PrepareUser(database.ModifyUser); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	sql, queries, err := db.Conn()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer sql.Close()

	repository := repositories.NewUserRepository(queries)

	if err = repository.UpdateUserById(uint32(userId), user); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	sql, queries, err := db.Conn()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer sql.Close()

	repository := repositories.NewUserRepository(queries)

	if err = repository.DeleteUserById(uint32(userId)); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
