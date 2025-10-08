package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWT(email string, role string) (string, error) {
	
	claims := jwt.MapClaims{
		"user_email": email,
		"user_role":  role,
		"exp":        time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}
	return t, nil
}

func ValidateJWT(cookie string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	return token, err
}
