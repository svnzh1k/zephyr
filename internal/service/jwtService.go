package service

import (
	"fmt"
	"zephyr-api-mod/internal/models"

	"github.com/dgrijalva/jwt-go"
)

var secretKey string = "hello"

func GenerateJWT(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to parse claims")
	}
	return claims, nil
}

func parseClaims(claims jwt.MapClaims) models.User {
	user := models.User{}

	if userID, ok := claims["userid"].(float64); ok {
		user.Id = int(userID)
	}
	if username, ok := claims["username"].(string); ok {
		user.Username = username
	}
	if role, ok := claims["role"].(string); ok {
		user.Role = role
	}
	return user
}
