package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func Config() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(w http.ResponseWriter, id, token string) error {
	data := map[string]string{
		"id":    id,
		"token": token,
	}

	encodedData, err := s.Encode("devbook-data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "devbook-data",
		Value:    encodedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("devbook-data")
	if err != nil {
		return nil, err
	}

	var values map[string]string

	if err = s.Decode("devbook-data", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}
