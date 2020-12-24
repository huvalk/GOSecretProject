package middleware

import (
	"GOSecretProject/core/auth/interfaces"
	"context"
	"github.com/kataras/golog"
	"net/http"
)

var userIDKey = "userID"

type AuthMiddlewareHandler struct {
	authRepository authInterfaces.AuthRepository
}

func NewAuthMiddlewareHandler(authRepository authInterfaces.AuthRepository) *AuthMiddlewareHandler {
	return &AuthMiddlewareHandler{authRepository: authRepository}
}

func (m *AuthMiddlewareHandler) UserRequired(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rID, ok := r.Context().Value("rID").(string)
		if !ok {
			rID = "no request id"
		}

		var userId uint64

		session, err := r.Cookie("session_id")
		if err != nil {
			golog.Infof("#%s: %s", rID, "No cookie")
			userId = 0
		} else {
			golog.Infof("#%s: %s", rID, session.Value)
			user, err := m.authRepository.CheckSession(session.Value)
			if err != nil {
				golog.Errorf("#%s: %s", rID, err.Error())
				userId = 0
			} else {
				golog.Infof("#%s: %s", rID, "success")
				userId = uint64(user.ID)
			}
		}

		r = r.WithContext(context.WithValue(r.Context(), userIDKey, userId))
		next.ServeHTTP(w, r)
	}
}
