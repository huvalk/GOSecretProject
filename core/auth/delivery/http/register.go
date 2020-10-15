package authHttp

import (
	authInterfaces "GOSecretProject/core/auth/interfaces"
	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoints(router *mux.Router, repository authInterfaces.AuthRepository) {
	h := NewHandler(repository)

	router.HandleFunc("/login", h.Login).Methods("POST")
	router.HandleFunc("/users", h.Login).Methods("POST")
}
