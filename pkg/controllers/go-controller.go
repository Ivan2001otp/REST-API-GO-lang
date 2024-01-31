package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"simple-api/app/pkg/Database"
	"simple-api/app/pkg/Models"
	"simple-api/app/pkg/constants"
)

func hello() string {
	return "hello"
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func sendError(w http.ResponseWriter, statusCode int, err string) {
	error_mess := map[string]interface{}{"error": err}
	sendResponse(w, statusCode, error_mess)
}

func createProdController(w http.ResponseWriter, req *http.Request) {
	var product Models.Product

	error := json.NewDecoder(req.Body).Decode(&product)

	if error != nil {
		fmt.Println("The error is ->", error)
		sendError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	result := Database.GetDB()

	err := product.CreateUser(result)

	if err != nil {
		sendError(w, http.StatusInternalServerError, "Went wrong while posting the Data")
		return
	}

	sendResponse(w, http.StatusOK, req)

}

func getAllProdController(w http.ResponseWriter, req *http.Request) {

	userCollection := Database.GetDB().Database(constants.DatabaseName).Collection(constants.ProductCollection)
	cursor, err := userCollection.Find(context.TODO(), map[string]interface{}{})

	if err != nil {
		fmt.Println("The error is ->", err)
		sendError(w, http.StatusInternalServerError, "gone wrong while fetching data")
		return
	}

	usersList := make([]Models.Product, 0)

	//this below code is executed at last.
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var temp Models.Product
		err := cursor.Decode(&cursor)
		if err != nil {
			sendError(w, http.StatusInternalServerError, "Gone wrong while iterating ")
			return
		}
		usersList = append(usersList, temp)
	}
	sendResponse(w, http.StatusOK, usersList)
}

func getProdByIdController(w http.ResponseWriter, req *http.Request) {

}

func updateProdController(w http.ResponseWriter, req *http.Request) {

}

func deleteProdController(w http.ResponseWriter, req *http.Request) {

}
