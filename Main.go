package main

import (
	"log"
	"net/http"

	"github.com/GuillaumeDSM/GoItemAPI/routing"
	"github.com/gorilla/handlers"
)

func main() {

	router := routing.NewRouter()
	corsObj := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsObj)(router)))

}
