// Writing the Business logic to connect the mongodb.
package main

import (
	/*
		"context"
		"fmt"
		"log"
		"go.mongodb.org/mongo-driver/bson"
		"go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/mongo/options"
	*/
	"fmt"
	"log"
	"net/http"
	"simple-api/app/pkg/Database"
	"simple-api/app/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	var mongo_url string = "mongodb://localhost:27017"
	res, err := Database.GetMongoClient(mongo_url)

	if err != nil {
		fmt.Println(err)
		return
	}

	if res != nil {
		fmt.Println("Connected to mongo db")
	}

	routes.RegisterUserRoutes(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8081", router))

}
