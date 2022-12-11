package midleware

import (
	"net/http"
	"strings"

	"github.com/deptech/internal/pkg/auth"
	response "github.com/deptech/internal/pkg/response"
)

func AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
		err := auth.ValidateToken(token)
		if err != nil {
			response.Response(w, response.View{
				Message:      "Error",
				ErrorMessage: err.Error(),
			}, http.StatusUnauthorized)

			return
		}

		next.ServeHTTP(w, r)
	})
}
