package main

import (
	"database/sql"
	"log"
	"net/http"
	"zephyr-api-mod/internal/handlers"
	"zephyr-api-mod/internal/middleware"
	"zephyr-api-mod/internal/service"
)

var Db *sql.DB

func main() {
	err := service.InitializeDatabase()
	if err != nil {
		log.Fatal("can't connect to the database", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /admin/register", middleware.TokenValidator((middleware.AdminRoleValidator(handlers.RegisterHandler))))
	mux.HandleFunc("POST /login", handlers.LoginHandler)

	http.ListenAndServe(":8080", middleware.RequestLogger(mux.ServeHTTP))
}
