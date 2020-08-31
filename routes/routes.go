package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sanketgupta07/go-rest-example/controllers"
	"github.com/sanketgupta07/go-rest-example/util/constants"
)

//Routers define all the endpoint route
func Routers(router *mux.Router) {
	router.HandleFunc(constants.EndpointHomeLink, controllers.HomeLink)
	router.HandleFunc(constants.EndpointCreateUser, controllers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc(constants.EndpointGetAllUsers, controllers.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc(constants.EndpointGetOneUser, controllers.GetOneUser).Methods(http.MethodGet)
	router.HandleFunc(constants.EndpointUpdateUser, controllers.UpdateUser).Methods(http.MethodPatch)
	router.HandleFunc(constants.EndpointDeleteUser, controllers.DeleteUser).Methods(http.MethodDelete)

	router.HandleFunc(constants.EndpointPokemon, controllers.PokeGetClient).Methods(http.MethodGet)
}
