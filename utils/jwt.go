package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shanomz7235/bookstore-back/models"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWT(user *models.LoginUser) (string, error) {

	claims := jwt.MapClaims{
		"user_email": user.Email,
		"user_role":  user.Role,
		"user_id":    user.ID,
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
