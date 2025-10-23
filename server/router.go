package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/chat", WsHandler)
	router.HandleFunc("/history", HistoryHandler)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(".")))

	return router
}
