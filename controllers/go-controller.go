package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ivan2001otp/REST-API-GO-lang/config"
	"github.com/Ivan2001otp/REST-API-GO-lang/constants"
	"github.com/Ivan2001otp/REST-API-GO-lang/models"
	"github.com/gorilla/mux"
)

// helper methods
func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func sendError(w http.ResponseWriter, statusCode int, err string) {
	err_mess := map[string]interface{}{"error": err}
	sendResponse(w, statusCode, err_mess)
}

//controller methods

func CreateUser(w http.ResponseWriter, req *http.Request) {
	var userModel models.User

	err := json.NewDecoder(req.Body).Decode(&userModel)

	if err != nil {
		fmt.Println("Something went wrong while posting")
		sendError(w, http.StatusBadRequest, "Invalid request payload")
	}

	err = userModel.ModelUser(config.GetDB())

	if err != nil {
		sendError(w, http.StatusInternalServerError, "Something went wrong-2")
		return
	}

	//returning the registered user.
	sendResponse(w, http.StatusAccepted, userModel)
}

func GetAllUser(w http.ResponseWriter, req *http.Request) {
	userCollection := config.GetDB().Database(constants.DBName).Collection(constants.UserCollection)

	//second arg is filter while fetching all the records from DB.
	cursor, err := userCollection.Find(context.TODO(), map[string]interface{}{})

	if err != nil {
		sendError(w, http.StatusInternalServerError, "get all user endpoint not working!")
		return
	}

	userList := make([]models.User, 0)
	defer cursor.Close(context.TODO()) //executed at last

	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)

		if err != nil {
			sendError(w, http.StatusInternalServerError, "something went wrong - 1")
			return
		}
		userList = append(userList, user)
	}
	sendResponse(w, http.StatusAccepted, userList)
}

func GetUserById(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)

	id := vars["id"]

	user, err := models.FetchUserById(config.GetDB(), id)

	if err != nil {
		sendError(w, http.StatusNotFound, err.Error())
		return
	}

	sendResponse(w, http.StatusOK, user)
}

func DeleteUserById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id := vars["id"]

	message, err := models.DeleteUserById(config.GetDB(), id)

	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, message)
}

func UpdateUserById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id := vars["id"]

	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	U, err := user.UpdateModelById(config.GetDB(), id)

	if err != nil {
		sendError(w, http.StatusNotFound, err.Error())
		return
	}

	sendResponse(w, http.StatusOK, U)

}
