package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"www.github.com/Ammce/banking-go-auth/handlers"
)

func NewAccountRoutes(r *mux.Router, uh handlers.UserHandlers) {
	r.HandleFunc("/", uh.Register).Methods(http.MethodPost)
}
