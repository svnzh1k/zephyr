package main

import (
	"database/sql"
	"log"
	"net/http"
	"zephyr-api-mod/internal/handlers"
	"zephyr-api-mod/internal/service"
	"zephyr-api-mod/middleware"
)

var Db *sql.DB

func main() {
	err := service.InitializeDatabase()
	if err != nil {
		log.Fatal("fatal", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", handlers.SignupHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	http.ListenAndServe(":8080", middleware.RequestLogger(mux.ServeHTTP))
}
