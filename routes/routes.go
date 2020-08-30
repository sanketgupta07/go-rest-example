package routes

import (
	"github.com/gorilla/mux"
	"github.com/sanketgupta07/go-rest-example/controllers"
)

//Routers define all the endpoint route
func Routers(router *mux.Router) {
	router.HandleFunc("/", controllers.HomeLink)
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{name}", controllers.GetOneUser).Methods("GET")
	router.HandleFunc("/users/{name}", controllers.UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{name}", controllers.DeleteUser).Methods("DELETE")
}
