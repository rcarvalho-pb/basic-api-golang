package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if _, err := cookies.Read(r); err != nil {
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
			return
		}
		next(w, r)
	}
}
