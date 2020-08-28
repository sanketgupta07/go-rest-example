package routes

import (
	"github.com/gorilla/mux"
	"github.com/sanketgupta07/go-rest-example/controllers"
)

func Routers(router *mux.Router) {
	router.HandleFunc("/", controllers.HomeLink)
	router.HandleFunc("/user", controllers.CreateEvent).Methods("POST")
	router.HandleFunc("/users", controllers.GetAllEvents).Methods("GET")
	router.HandleFunc("/users/{name}", controllers.GetOneEvent).Methods("GET")
	router.HandleFunc("/users/{name}", controllers.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/users/{name}", controllers.DeleteEvent).Methods("DELETE")
}
