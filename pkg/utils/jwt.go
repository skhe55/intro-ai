package utils

import (
	"intro-ai/config"
	"intro-ai/internal/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserName string `json:"username"`
	Login    string `json:"login"`
	ID       uint64 `json:"id"`
	jwt.StandardClaims
}

func GenerateJWT(user *models.User, config *config.Config) (string, string, error) {
	claims := &Claims{
		UserName: user.UserName,
		Login:    user.Login,
		ID:       user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.JwtSecretKey))
	if err != nil {
		return "", "", err
	}

	return tokenString, time.Now().Add(time.Minute * 60).UTC().String(), nil
}
