package middleware

import (
	"fmt"
	"learningGorillamux/data"
	"net/http"
)

func ValidateUser(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("username")
		password := r.Header.Get("password")

		if data.Users[username] != password || username == "" {
			w.Write([]byte("Failed to authenticate"))
			return
		}
		f(w, r)

	}
}

func ValidateOwner(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("username")
		password := r.Header.Get("password")
		IfOwner := r.Header.Get("userType")
		if IfOwner != "owner" {
			w.Write([]byte("you are not owner"))
			return
		}
		if data.Users[username] != password || username == "" {
			w.Write([]byte("Failed to authenticate"))
			return
		}
		f(w, r)

	}
}

func TrackNumberOfRequests(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data.NubmerOfRequests = data.NubmerOfRequests + 1
		fmt.Println("Request Number : ", data.NubmerOfRequests)

		f.ServeHTTP(w, r)
	})
}
