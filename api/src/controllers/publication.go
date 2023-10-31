package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	author_id, err := authentication.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	var publication models.Publication
	if err = json.Unmarshal(requestBody, &publication); err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthorID = int32(author_id)

	db, queries, err := db.Conn()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	publicationRepository := repositories.NewPublicationRepository(queries)

	if err = publication.Prepare(); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	publicationId, err := publicationRepository.CreatePublication(publication)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	publication.ID = int32(publicationId)

	response.JSON(w, http.StatusCreated, publication)
}

func FindPublication(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, queries, err := db.Conn()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationRepository(queries)

	publications, err := repository.FindPublications(userId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publications)
}

func FindPublicationById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationId, err := strconv.ParseInt(params["publicationId"], 10, 32)
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

	repository := repositories.NewPublicationRepository(queries)

	publication, err := repository.FindPublicationById(publicationId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publication)
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)

	publicationId, err := strconv.ParseInt(params["publicationId"], 10, 32)
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

	repository := repositories.NewPublicationRepository(queries)

	publication, err := repository.FindPublicationById(publicationId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if userId != int64(publication.AuthorID) {
		response.ERROR(w, http.StatusForbidden, errors.New("cant update other person publication"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var newPublication models.Publication

	if err = json.Unmarshal(requestBody, &newPublication); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if err = repository.UpdatePublication(publicationId, newPublication); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusAccepted, nil)
}

func DeletePublication(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)

	publicationId, err := strconv.ParseInt(params["publicationId"], 10, 32)
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

	repository := repositories.NewPublicationRepository(queries)

	publication, err := repository.FindPublicationById(publicationId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if publication.AuthorID != int32(userId) {
		response.ERROR(w, http.StatusForbidden, errors.New("cant delete other user publication"))
		return
	}

	if err = repository.DeletePublicationById(publicationId); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusAccepted, nil)
}

func GetUserPublications(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseInt(params["userId"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, queries, err := db.Conn()
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationRepository(queries)

	publications, err := repository.GetUserPublications(userId)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, publications)
}

func LikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationId, err := strconv.ParseInt(params["publicationId"], 10, 32)
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

	repository := repositories.NewPublicationRepository(queries)

	if err = repository.LikePublication(publicationId); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusAccepted, nil)
}

func DislikePublication(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicationId, err := strconv.ParseInt(params["publicationId"], 10, 32)
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

	repository := repositories.NewPublicationRepository(queries)

	if err = repository.DislikePublication(publicationId); err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusAccepted, nil)
}
