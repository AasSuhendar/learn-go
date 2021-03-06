package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dimaskiddo/simple-go/api-mongo/helpers"
	"github.com/dimaskiddo/simple-go/api-mongo/models"
	"github.com/dimaskiddo/simple-go/api-mongo/routers"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// Function to Get User
func GetUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	var response routers.ResponseGetUser

	// Database Query
	err := helpers.DB.C("users").Find(bson.M{}).All(&users)
	if err != nil {
		log.Print(err)
	}

	// Set Response Data
	response.Status = true
	response.Code = http.StatusOK
	response.Message = "Success"
	response.Data = users

	// Write Response Data to HTTP
	routers.ResponseWrite(w, response.Code, response)
}

// Function to Get User By ID
func GetUserById(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Handle Error If Parameters ID is Empty
	if len(params["id"]) == 0 {
		routers.ResponseBadRequest(w)
	} else {
		var user models.User
		var users []models.User
		var response routers.ResponseGetUser

		// Database Query
		err := helpers.DB.C("users").FindId(bson.ObjectIdHex(params["id"])).One(&user)
		if err != nil {
			log.Print(err)
		}

		// Convert Selected User from Users Array to Single User Array
		users = append(users, user)

		// Set Response Data
		response.Status = true
		response.Code = http.StatusOK
		response.Message = "Success"
		response.Data = users

		// Write Response Data to HTTP
		routers.ResponseWrite(w, response.Code, response)
	}
}

// Function to Add User
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var response routers.Response

	// Decode JSON from Request Body to User Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Set User ID to New Generated ID
	user.ID = bson.NewObjectId()

	// Database Query
	err := helpers.DB.C("users").Insert(&user)
	if err != nil {
		log.Print(err)
	}

	// Set Response Data
	response.Status = true
	response.Code = http.StatusCreated
	response.Message = "Success"

	// Write Response Data to HTTP
	routers.ResponseWrite(w, response.Code, response)
}

// Function to Update User By ID
func PutUserById(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Handle Error If Parameters ID is Empty
	if len(params["id"]) == 0 {
		routers.ResponseBadRequest(w)
	} else {
		var user models.User
		var response routers.Response

		// Decode JSON from Request Body to User Data
		// Use _ As Temporary Variable
		_ = json.NewDecoder(r.Body).Decode(&user)

		// Database Query
		err := helpers.DB.C("users").UpdateId(params["id"], &user)
		if err != nil {
			log.Print(err)
		}

		// Set Response Data
		response.Status = true
		response.Code = http.StatusCreated
		response.Message = "Success"

		// Write Response Data to HTTP
		routers.ResponseWrite(w, response.Code, response)
	}
}

// Function to Delete User By ID
func DelUserById(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Handle Error If Parameters ID is Empty
	if len(params["id"]) == 0 {
		routers.ResponseBadRequest(w)
	} else {
		var user models.User
		var response routers.Response

		// Database Query Get User By ID
		err := helpers.DB.C("users").FindId(bson.ObjectIdHex(params["id"])).One(&user)
		if err != nil {
			log.Print(err)
		}

		// Database Query Delete User
		err = helpers.DB.C("users").Remove(&user)
		if err != nil {
			log.Print(err)
		}

		// Set Response Data
		response.Status = true
		response.Code = http.StatusOK
		response.Message = "Success"

		// Write Response Data to HTTP
		routers.ResponseWrite(w, response.Code, response)
	}
}
