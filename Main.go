package main

import (
	"log"
	"net/http"

	"github.com/GuillaumeDSM/GoItemAPI/routing"
)

func main() {

	router := routing.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
