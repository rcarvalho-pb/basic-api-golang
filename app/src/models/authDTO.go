package models

type AuthDTO struct {
	ID    uint32 `json:"user_id"`
	Token string `json:"token"`
}
