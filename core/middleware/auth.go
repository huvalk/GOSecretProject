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

func (m *AuthMiddlewareHandler) UserRequired(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rID, ok := r.Context().Value("rID").(string)
		if !ok {
			rID = "no request id"
		}

		session, err := r.Cookie("session_id")
		if err != nil {
			golog.Infof("#%s: %s", rID, "No cookie")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		golog.Infof("#%s: %s", rID, session.Value)
		userID, err := m.authRepository.CheckSession(session.Value)
		if err != nil {
			golog.Errorf("#%s: %s", rID, err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			golog.Infof("#%s: %s", rID, "success")
		}

		r = r.WithContext(context.WithValue(r.Context(), userIDKey, userID))
		next.ServeHTTP(w, r)
	}
}
