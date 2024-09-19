package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// func RequestLogger(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("Received request: %s %s at %s\n", r.Method, r.URL.Path, time.Now().Format(time.DateTime))
// 		next(w, r)
// 	}
// }

// const (
// 	Reset  = "\033[0m"
// 	Red    = "\033[31m"
// 	Green  = "\033[32m"
// 	Yellow = "\033[33m"
// 	Blue   = "\033[34m"
// )

// func RequestLogger(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Определяем цвет в зависимости от метода запроса
// 		var color string
// 		switch r.Method {
// 		case http.MethodGet:
// 			color = Green
// 		case http.MethodPost:
// 			color = Yellow
// 		case http.MethodPatch:
// 			color = Blue
// 		case http.MethodDelete:
// 			color = Red
// 		default:
// 			color = Reset
// 		}

// 		// Логируем информацию о запросе с цветом
// 		fmt.Printf("%s[%s] Method: %s | URL: %s%s\n",
// 			color,
// 			time.Now().Format("2006-01-02 15:04:05"),
// 			r.Method,
// 			r.URL.Path,
// 			Reset, // Сброс цвета после вывода
// 		)

// 		// Вызов следующего обработчика
// 		next(w, r)
// 	}
// }

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

// RequestLogger логирует запросы с цветами для методов
func RequestLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Определяем цвет в зависимости от метода запроса
		var color string
		switch r.Method {
		case http.MethodGet:
			color = Blue
		case http.MethodPost:
			color = Green
		case http.MethodPut:
			color = Yellow
		case http.MethodDelete:
			color = Red
		default:
			color = Reset
		}

		// Логируем запрос в одну строку
		fmt.Printf("[%s] %sMethod: %s%s | URL: %s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			color,
			r.Method,
			Reset, // Сброс цвета
			r.URL.Path,
		)

		// Вызов следующего обработчика
		next(w, r)
	}
}
