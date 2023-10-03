package utils

import (
	"errors"
	"html"
	"intro-ai/config"
	"intro-ai/internal/models"
	"intro-ai/pkg/utils/httpError"
	"net/http"
	"strings"
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

func ExtractJWTFromRequest(r *http.Request) (map[string]interface{}, error) {
	tokenString := ExtractBearerToken(r)

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (jwtKey interface{}, err error) {
		return jwtKey, err
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, httpError.InvalidJWTSignature
		}
		if err.Error() == "key is of invalid type" {
			return claims, nil
		}
		return nil, err
	}
	if !token.Valid {
		return nil, httpError.InvalidJWTToken
	}
	return claims, nil
}

func ExtractBearerToken(r *http.Request) string {
	headerAuthorization := r.Header.Get("Authorization")
	bearerToken := strings.Split(headerAuthorization, " ")
	return html.EscapeString(bearerToken[1])
}
