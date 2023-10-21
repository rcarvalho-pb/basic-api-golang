package controllers

import (
	"api/db/database"
	"api/src/authentication"
	"api/src/db"
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

	id, err := authentication.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusBadGateway, err)
		return
	}

	if id == uint32(userId) {
		response.ERROR(w, http.StatusForbidden, errors.New("cant update different user"))
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

	id, err := authentication.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusBadGateway, err)
		return
	}

	if id == uint32(userId) {
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

	if err = repository.DeleteUserById(uint32(userId)); err != nil {
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
	followedId, err := strconv.ParseUint(params["userId"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if id == uint32(followedId) {
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

	if err = repository.FollowUser(id, uint32(followedId)); err != nil {
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
	followedId, err := strconv.ParseUint(params["userId"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if id == uint32(followedId) {
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

	if err = repository.UnfollowUser(id, uint32(followedId)); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusAccepted, nil)
}

func UserFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 32)
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

	users, err := repository.GetUsersFollows(uint32(userId))
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusFound, users)
}

func UserFollowed(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 32)
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

	users, err := repository.GetUserFollowed(uint32(userId))
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
	userId, err := strconv.ParseUint(params["userId"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if userId != uint64(id) {
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

	user, err := repository.GetUserById(int32(userId))
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

	if err = repository.UpdatePassword(uint32(userId), string(hashedPassword)); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, nil)
}
