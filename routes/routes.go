package routes

import (
	"github.com/gorilla/mux"
	"github.com/sanketgupta07/go-rest-example/controllers"
)

func Routers(router *mux.Router) {
	router.HandleFunc("/", controllers.HomeLink)
	router.HandleFunc("/event", controllers.CreateEvent).Methods("POST")
	router.HandleFunc("/events", controllers.GetAllEvents).Methods("GET")
	router.HandleFunc("/events/{name}", controllers.GetOneEvent).Methods("GET")
	router.HandleFunc("/events/{name}", controllers.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/events/{name}", controllers.DeleteEvent).Methods("DELETE")
}
