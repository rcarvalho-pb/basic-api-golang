package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io"
	"net/http"
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
	userRepository := repositories.NewUserRepository(queries)
	
	user, err := userRepository.GetUserById(int32(author_id))
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthorNick = user.Nick

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
	// userId, err := authentication.ExtractUserId(r)
	// if err != nil {
	// 	response.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }


}

func FindPublicationById(w http.ResponseWriter, r *http.Request) {
	
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {

}

func DeletePublication(w http.ResponseWriter, r *http.Request) {

}