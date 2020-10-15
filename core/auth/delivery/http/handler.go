package authHttp

import (
	authInterfaces "GOSecretProject/core/auth/interfaces"
	"GOSecretProject/core/model/base"
	"encoding/json"
	"github.com/kataras/golog"
	"net/http"
)

type Handler struct {
	repo authInterfaces.AuthRepository
}

func NewHandler(repo authInterfaces.AuthRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {
	golog.Infof("Test ")

	w.Write([]byte("hello"))
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user base.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		golog.Errorf("Register error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	golog.Infof("Register: %s", user.Login)

	err = h.repo.Register(user)

	if err == nil {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte{})
	} else {
		golog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var user base.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		golog.Errorf("Login error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	golog.Infof("Login: %s", user)

	var code int
	user.ID, user.Session, code, err = h.repo.Login(user)

	switch code {
	case 201:
		w.WriteHeader(http.StatusCreated)
		json, _ := json.Marshal(user)
		w.Write(json)
	default:
		golog.Error(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte{})
	}
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	var user base.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		golog.Errorf("Logout error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	golog.Infof("Logout: %s", user)

	err = h.repo.Logout(user.Session)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte{})
	} else {
		golog.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
	}
}
