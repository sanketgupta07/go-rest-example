package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sanketgupta07/go-rest-example/routes"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	routes.Routers(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
