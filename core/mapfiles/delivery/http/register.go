package mapFilesHttp

import (
	"github.com/gorilla/mux"
	)

func RegisterHTTPEndpoints(router *mux.Router) {
	h := NewHandler()

	router.HandleFunc("/map_files", h.GetMapFiles).Methods("GET")
}
