package router

import (
	"github.com/gorilla/mux"
	"main/internal/handlers"
	"net/http"
)

func NewRouter(h *handlers.Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/note", h.Create).Methods(http.MethodPost)
	router.HandleFunc("/note/{id}", h.Read).Methods(http.MethodGet)
	router.HandleFunc("/note/{id}", h.Update).Methods(http.MethodPatch)
	router.HandleFunc("/note/{id}", h.Delete).Methods(http.MethodDelete)
	return router
}
