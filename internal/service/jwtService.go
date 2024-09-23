package service

import (
	"errors"
	"time"
	"zephyr-api-mod/internal/models"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey string = "hello"
var tokenExpiration = time.Hour * 24

func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(tokenExpiration).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(secretKey))
}

func ValidateJWT(token string) (*models.User, error) {
	// Parse the token and extract claims
	claims := &jwt.MapClaims{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// Check that the signing method is what we expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		// Return the secret key for verification
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err // Return the error if parsing fails
	}

	if !t.Valid {
		return nil, errors.New("invalid token") // Return an error if the token is not valid
	}

	// Extract user information from claims
	user := &models.User{
		Id:       int((*claims)["id"].(float64)), // Convert float64 to int
		Username: (*claims)["username"].(string),
		Role:     (*claims)["role"].(string),
	}

	return user, nil // Return the user if everything is fine
}
