package middlewares

import (
	"net/http"
	"strings"

	"github.com/aleksanderpalamar/rbac-api/auth"
	"github.com/dgrijalva/jwt-go"
)

func IsAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil {
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(r.Header["Authorization"][0], " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return auth.SigningKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
