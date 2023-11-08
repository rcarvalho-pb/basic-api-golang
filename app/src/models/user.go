package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/request"
)

type User struct {
	ID           int32         `json:"ID,omitempty"`
	Name         string        `json:"Name,omitempty"`
	Nick         string        `json:"Nick,omitempty"`
	Email        string        `json:"Email,omitempty"`
	Password     string        `json:"Password,omitempty"`
	CreatedAt    sql.NullTime  `json:"CreatedAt,omitempty"`
	Followers    []User        `json:"Followers,omitempty"`
	Follows      []User        `json:"Follows,omitempty"`
	Publications []Publication `json:"Publications,omitempty"`
}

func GetCompleteUser(userId int64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followsChannel := make(chan []User)
	publicationsChannel := make(chan []Publication)

	go GetUserData(userChannel, userId, r)
	go GetFollowers(followersChannel, userId, r)
	go GetFollows(followsChannel, userId, r)
	go GetPublications(publicationsChannel, userId, r)

	var (
		user         User
		followers    []User
		follows      []User
		publications []Publication
	)

	for i := 0; i < 4; i++ {
		select {
		case loadUser := <-userChannel:
			if loadUser.ID == 0 {
				return User{}, errors.New("error while getting user")
			}
			user = loadUser

		case loadFollowers := <-followersChannel:
			if loadFollowers == nil {
				return User{}, errors.New("error while getting followers")
			}
			followers = loadFollowers

		case loadFollows := <-followsChannel:
			if loadFollows == nil {
				return User{}, errors.New("error while getting follows")
			}
			follows = loadFollows

		case loadPublications := <-publicationsChannel:
			if loadPublications == nil {
				return User{}, errors.New("error while getting publications")
			}
			publications = loadPublications
		}
	}

	user.Followers = followers
	user.Follows = follows
	user.Publications = publications

	return user, nil
}

func GetUserData(channel chan<- User, userId int64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.ApiUrl, userId)
	res, err := request.Request(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer res.Body.Close()

	var user User
	if err = json.NewDecoder(res.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

func GetFollowers(channel chan<- []User, userId int64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.ApiUrl, userId)
	res, err := request.Request(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer res.Body.Close()

	var followers []User
	if err = json.NewDecoder(res.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

func GetFollows(channel chan<- []User, userId int64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followed", config.ApiUrl, userId)
	res, err := request.Request(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer res.Body.Close()

	var follows []User
	if err = json.NewDecoder(res.Body).Decode(&follows); err != nil {
		channel <- nil
		return
	}

	if follows == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- follows
}

func GetPublications(channel chan<- []Publication, userId int64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/publications", config.ApiUrl, userId)
	res, err := request.Request(r, http.MethodGet, url, nil)
	if err != nil {
		log.Println("Error request")
		channel <- nil
		return
	}
	defer res.Body.Close()

	var publications []Publication
	if err = json.NewDecoder(res.Body).Decode(&publications); err != nil {
		channel <- nil
		return
	}

	if publications == nil {
		channel <- make([]Publication, 0)
		return
	}

	channel <- publications
}
