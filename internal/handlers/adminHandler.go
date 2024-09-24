package handlers

import (
	"encoding/json"
	"net/http"
	"zephyr-api-mod/internal/models"
	"zephyr-api-mod/internal/service"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)
	if user.Username == "" || user.Code < 999 || user.Code > 9999 || user.Password == "" || len(user.Password) > 30 || len(user.Password) < 3 {
		http.Error(w, "a property for registration is missing", http.StatusBadRequest)
		return
	}
	err := service.AddUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Successfully added a new user with role = waiter"))
}
