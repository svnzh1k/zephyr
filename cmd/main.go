package main

import (
	"net/http"
	"zephyr-api-mod/middleware"
	// "zephyr-api-mod/middleware"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/seeyou", anotherFunc)
	http.ListenAndServe(":8080", middleware.RequestLogger(mux.ServeHTTP))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func anotherFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Another function"))
}
