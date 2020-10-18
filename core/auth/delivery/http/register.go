package authHttp

import (
	authInterfaces "GOSecretProject/core/auth/interfaces"
	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoints(router *mux.Router, repository authInterfaces.AuthRepository) {
	h := NewHandler(repository)

	router.HandleFunc("/login", h.Login).Methods("POST")
	router.HandleFunc("/users", h.Register).Methods("POST")
	router.HandleFunc("/phone/{phone:[0-9]+}/confirm", h.Confirm).Methods("GET")
	router.HandleFunc("/session", h.CheckSession).Methods("POST")
	router.HandleFunc("/test", h.Test).Methods("GET")
}
