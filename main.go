package main

import (
	"log"
	"net/http"

	"github.com/sanketgupta07/go-rest-example/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	routes.Routers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
