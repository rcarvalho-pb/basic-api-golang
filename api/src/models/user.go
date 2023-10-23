package models

import (
	"api/db/database"
	"api/src/security"
	"errors"
	"regexp"
	"strings"
)

// type User struct {
// 	ID        int32        `json:"id,omitempty"`
// 	Name      string       `json:"name,omitempty"`
// 	Nick      string       `json:"nick,omitempty"`
// 	Email     string       `json:"email,omitempty"`
// 	Password  string       `json:"password,omitempty"`
// 	CreatedAt sql.NullTime `json:"created_at,omitempty"`
// }

type Type uint32

const (
	NewUser = iota
	ModifyUser
)

type User database.User

func (u *User) PrepareUser(stage Type) error {
	if err := u.validateUser(stage); err != nil {
		return err
	}

	if err := u.formatUser(stage); err != nil {
		return err
	}
	return nil
}

func (u *User) validateUser(stage Type) error {
	if u.Name == "" {
		return errors.New("name field is mandatory")
	}
	if u.Nick == "" {
		return errors.New("nick field is mandatory")
	}
	nickRegexp := regexp.MustCompile(`^([\w\d\?\!\_\-\+]+)$`)
	if !nickRegexp.MatchString(u.Nick) {
		return errors.New("inválid nick format")
	}

	if u.Email == "" {
		return errors.New("email field is mandatory")
	}
	emailRegexp := regexp.MustCompile(`^([\w\d\.\-\!\?]+)@([\w\d]{2,}).([\w\d]{2,}).([\w\d]{2,})+?$`)
	if !emailRegexp.MatchString(u.Email) {
		return errors.New("inválid email format")
	}

	if stage == NewUser && u.Password == "" {
		return errors.New("password field is mandatory")
	}
	return nil
}

func (u *User) formatUser(stage Type) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
	u.Password = strings.TrimSpace(u.Password)

	if stage == NewUser {
		hashedPassword, err := security.Hash(u.Password)
		if err != nil {
			return err
		}

		u.Password = string(hashedPassword)
	}

	return nil
}
