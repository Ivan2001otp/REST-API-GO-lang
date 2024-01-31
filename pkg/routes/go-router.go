package routes

import (
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {

	//create the user route-post
	router.HandleFunc("/user", controller.createUserController).Methods("POST")

	//get All user route
	router.HandleFunc("/users", controller.getAllUserController).Methods("GET")

	//get user by ID
	router.HandleFunc("/user/{id}", controller.getUserByIdController).Methods("GET")

	//update user by ID route.
	router.HandleFunc("/user/{id}", controller.updateUserController).Methods("PUT")

	//delete user route
	router.HandleFunc("/user/{id}", controller.deleteUserController).Methods("DELETE")

}
