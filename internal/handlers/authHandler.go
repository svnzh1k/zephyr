package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"zephyr-api-mod/internal/service"

	"golang.org/x/crypto/bcrypt"
)

func respond(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	w.Write([]byte(msg))
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header.Get("Authorization"), " ")[1]
	fmt.Print(token)
	fmt.Println(service.ValidateJWT(token))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		respond(w, err.Error(), http.StatusInternalServerError)
		return
	}
	username, exists := data["username"]
	if !exists {
		respond(w, "missing username property", http.StatusNotAcceptable)
		return
	}
	password, exists := data["password"]
	if !exists {
		respond(w, "missing password property", http.StatusNotAcceptable)
		return
	}
	user, err := service.GetUserByUsername(username)
	if err != nil {
		respond(w, "No such user", http.StatusNotAcceptable)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		respond(w, "password is incorrect", http.StatusBadRequest)
		return
	}

	token, err := service.GenerateJWT(user)
	if err != nil {
		respond(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Authorization", "Bearer "+token)
	respond(w, "Login successful", 200)
}
