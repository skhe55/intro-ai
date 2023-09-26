package middleware

import (
	"context"
	"fmt"
	"intro-ai/pkg/utils/httpError"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var (
	UNAUTHORIZED_MESSAGE = "Unauthorized user"
)

func (mw *MiddlewareManager) AuthJWTMiddleware() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			bearerHeader := r.Header.Get("Authorization")
			httpError := httpError.NewHttpError(w)

			if bearerHeader != "" {
				headerParts := strings.Split(bearerHeader, " ")

				if len(headerParts) != 2 {
					mw.logger.Error("Failing while parse auth header")
					httpError.NonInternalError(http.StatusUnauthorized, UNAUTHORIZED_MESSAGE)
					return
				}

				tokenString := headerParts[1]
				if err := mw.validateJWTToken(tokenString, r.Context()); err != nil {
					mw.logger.Error("Not valid jwt token")
					httpError.NonInternalError(http.StatusUnauthorized, UNAUTHORIZED_MESSAGE)
					return
				}

				next(w, r)
			} else {
				httpError.NonInternalError(http.StatusUnauthorized, UNAUTHORIZED_MESSAGE)
				return
			}
		}
	}
}

func (mw *MiddlewareManager) validateJWTToken(tokenString string, ctx context.Context) error {
	if tokenString == "" {
		return httpError.InvalidJWTToken
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method %v", token.Header["alg"])
		}
		secret := []byte(mw.cfg.JwtSecretKey)
		return secret, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return httpError.InvalidJWTToken
	}

	return nil
}
