package main

import (
	"log"
	"net/http"

	"github.com/GuillaumeDSM/GoItemAPI/routing"
	"github.com/gorilla/handlers"
)

func main() {

	router := routing.NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST"})
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(headersOk, originsOK, methodsOk)(router)))

}
