package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	user "www.github.com/Ammce/banking-go-auth/adapters/User"
	"www.github.com/Ammce/banking-go-auth/config"
	users "www.github.com/Ammce/banking-go-auth/domain/Users"
	"www.github.com/Ammce/banking-go-auth/handlers"
	"www.github.com/Ammce/banking-go-auth/routes"
)

func StartServer() {
	r := mux.NewRouter()

	db := config.ConfigureMongoDB()

	userRepositoryDB := user.NewUserRepositoryDB(db)
	userService := users.NewUserService(userRepositoryDB)
	userHandlers := handlers.NewUserHandlers(*userService)

	usersRouter := r.PathPrefix("/users").Subrouter()
	routes.NewUserRoutes(usersRouter, userHandlers)

	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
