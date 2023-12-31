package controllers

import (
	"encoding/json"
	"fmt"
	"learningGorillamux/data"
	"learningGorillamux/datatypes"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user datatypes.UserCredentials

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while decoding json request body", http.StatusBadRequest)
		return
	}

	_, exists := data.Users[user.Username]

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user doesn't exist"))
		return
	}

	if data.Users[user.Username] == user.Password {
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("Valid user"))
		return
	}

	w.Write([]byte("Invalid password"))

}

func Signup(w http.ResponseWriter, r *http.Request) {
	var newUser datatypes.UserCredentials

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Error while decoding json request body", http.StatusBadRequest)
		return
	}

	_, exists := data.Users[newUser.Username]
	if exists {
		w.Write([]byte("User already exists!"))
		return
	}

	data.Users[newUser.Username] = newUser.Password
	w.Write([]byte("New user created"))
}
