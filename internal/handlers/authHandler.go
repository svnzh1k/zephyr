package handlers

import (
	"encoding/json"
	"net/http"
	"zephyr-api-mod/internal/service"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	username, exists := data["username"]
	if !exists {
		http.Error(w, "missing username property", http.StatusNotAcceptable)
		return
	}
	password, exists := data["password"]
	if !exists {
		http.Error(w, "missing password property", http.StatusNotAcceptable)
		return
	}
	user, err := service.GetUserByUsername(username)
	if err != nil {
		http.Error(w, "No such user", http.StatusNotAcceptable)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, "password is incorrect", http.StatusBadRequest)
		return
	}

	token, err := service.GenerateJWT(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Authorization", "Bearer "+token)
	w.Write([]byte("Login successful"))
}
