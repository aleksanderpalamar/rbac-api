package middlewares

import (
	"net/http"
)

func IsAuthorized(role string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Aqui, você implementaria a lógica para verificar o papel do usuário
		// Por exemplo, extrair o papel do usuário de um token JWT ou sessão
		userRole := "admin" // Isso deve ser substituído pela lógica real

		if userRole != role {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next(w, r)
	})
}
