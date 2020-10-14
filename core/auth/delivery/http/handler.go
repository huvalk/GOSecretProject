package authHttp

import (
	"GOSecretProject/core/model/base"
	"encoding/json"
	"github.com/kataras/golog"
	"net/http"
	"time"
)

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{

	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var user base.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		golog.Errorf("Login ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Login == "root" && user.Password == "root" {
		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    user.Login + user.Password,
			Expires:  time.Now().Add(time.Hour),
			MaxAge:   100000,
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)

		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		//json := byte("")
		//w.Write(json)
		w.Write([]byte{})
	}
}