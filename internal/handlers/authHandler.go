package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zephyr-api-mod/internal/service"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Write([]byte("Error mapping json, please check the LoginHandler"))
		return
	}
	username, exists := data["username"]
	if !exists {
		w.Write([]byte("missing username property"))
		return
	}
	// password, exists := data["password"]
	if !exists {
		w.Write([]byte("missing password property"))
		return
	}
	user := service.GetUserByUsername(username)
	fmt.Print(user)

}
