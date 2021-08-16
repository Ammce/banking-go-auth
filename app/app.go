package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()

	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
