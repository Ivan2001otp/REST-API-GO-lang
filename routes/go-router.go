package routes

import (
	"github.com/Ivan2001otp/REST-API-GO-lang/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetAllUser).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.DeleteUserById).Methods("DELETE")

}
