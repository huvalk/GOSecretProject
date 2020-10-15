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

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user base.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		golog.Errorf("Register ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	golog.Infof("Register ", user)

	err = h.repo.Register(user)

	if err != nil {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte{})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
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
	golog.Infof("Login ", user)

	var code int
	user.ID, user.Session, code = h.repo.Login(user)

	switch code {
	case 201:
		w.WriteHeader(http.StatusCreated)
		json, _ := json.Marshal(user)
		w.Write(json)
	default:
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte{})
	}
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	var user base.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		golog.Errorf("Logout ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	golog.Infof("Logout ", user)

	err = h.repo.Logout(user.Session)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte{})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte{})
	}
}
