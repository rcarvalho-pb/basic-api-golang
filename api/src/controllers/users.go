package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"api/src/security"
	"encoding/json"
	"errors"
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

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = user.PrepareUser(models.NewUser); err != nil {
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
	userId, err := strconv.ParseInt(params["userId"], 10, 64)
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

	user, err := repository.GetUserById(userId)
	if err != nil {
		response.ERROR(w, http.StatusNotFound, err)
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseInt(params["userId"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	id, err := authentication.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusBadGateway, err)
		return
	}

	if id != userId {
		response.ERROR(w, http.StatusForbidden, errors.New("cant update different user"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = user.PrepareUser(models.ModifyUser); err != nil {
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

	if err = repository.UpdateUserById(userId, user); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseInt(params["userId"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	id, err := authentication.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusBadGateway, err)
		return
	}

	if id == userId {
		response.ERROR(w, http.StatusForbidden, errors.New("cant delete different user"))
		return
	}

	sql, queries, err := db.Conn()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer sql.Close()

	repository := repositories.NewUserRepository(queries)

	if err = repository.DeleteUserById(userId); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	id, err := authentication.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	followedId, err := strconv.ParseInt(params["userId"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if id == followedId {
		response.ERROR(w, http.StatusForbidden, errors.New("cant follow yourself"))
		return
	}

	db, queries, err := db.Conn()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(queries)

	if err = repository.FollowUser(id, followedId); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	id, err := authentication.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	followedId, err := strconv.ParseInt(params["userId"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if id == followedId {
		response.ERROR(w, http.StatusForbidden, errors.New("cant unfollow yourself"))
		return
	}

	db, queries, err := db.Conn()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(queries)

	if err = repository.UnfollowUser(id, followedId); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusAccepted, nil)
}

func UserFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseInt(params["userId"], 10, 32)
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

	users, err := repository.GetUsersFollows(userId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusFound, users)
}

func UserFollowed(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseInt(params["userId"], 10, 32)
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

	users, err := repository.GetUserFollowed(userId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusFound, users)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	id, err := authentication.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	userId, err := strconv.ParseInt(params["userId"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if userId != id {
		response.ERROR(w, http.StatusForbidden, errors.New("cant update another users password"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var passwords struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err = json.Unmarshal(requestBody, &passwords); err != nil {
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

	user, err := repository.GetUserById(userId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if security.ValidatePassword(passwords.OldPassword, user.Password) != nil {
		response.ERROR(w, http.StatusConflict, errors.New("wrong password"))
		return
	}

	hashedPassword, err := security.Hash(passwords.NewPassword)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
	}

	if err = repository.UpdatePassword(userId, string(hashedPassword)); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}
