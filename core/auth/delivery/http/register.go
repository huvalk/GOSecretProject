package authHttp

import (
	"github.com/gorilla/mux"
	)

func RegisterHTTPEndpoints(router *mux.Router) {
	h := NewHandler()

	router.HandleFunc("/login", h.Login).Methods("POST")
}
