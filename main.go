package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ivan2001otp/REST-API-GO-lang/config"
	"github.com/Ivan2001otp/REST-API-GO-lang/routes"
	"github.com/gorilla/mux"
)

func main() {
	config.ConnectToDB() //initializes the connection to database.

	router := mux.NewRouter()

	routes.RegisterRoutes(router)

	http.Handle("/", router)
	fmt.Println("listening to port 8099")
	log.Fatal(http.ListenAndServe(":8099", router))

}
