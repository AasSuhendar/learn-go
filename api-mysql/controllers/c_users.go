package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dimaskiddo/simple-go/api-mysql/helpers"
	"github.com/dimaskiddo/simple-go/api-mysql/models"
	"github.com/dimaskiddo/simple-go/api-mysql/routers"

	"github.com/gorilla/mux"
)

// Function to Get User
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var users []models.User
	var response routers.ResponseGetUser

	// Database Query
	rows, err := helpers.DB.Query("SELECT * FROM users")
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	// Populate Data
	for rows.Next() {
		// Match / Binding Database Field with Struct
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Fatal(err.Error())
		} else {
			// Append User Struct to Users Array of Struct
			users = append(users, user)
		}
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
		// Get ID Parameters From URI Then Convert it to Integer
		userID, err := strconv.Atoi(params["id"])
		if err == nil {
			var user models.User
			var users []models.User
			var response routers.ResponseGetUser

			// Database Query
			rows, err := helpers.DB.Query("SELECT * FROM users WHERE id=? LIMIT 1", userID)
			if err != nil {
				log.Print(err)
			}
			defer rows.Close()

			// Populate Data
			for rows.Next() {
				// Match / Binding Database Field with Struct
				if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
					log.Fatal(err.Error())
				} else {
					// Append User Struct to Users Array of Struct
					users = append(users, user)
				}
			}

			// Set Response Data
			response.Status = true
			response.Code = http.StatusOK
			response.Message = "Success"
			response.Data = users

			// Write Response Data to HTTP
			routers.ResponseWrite(w, response.Code, response)
		} else {
			routers.ResponseInternalError(w)
		}
	}
}

// Function to Add User
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var response routers.Response

	// Decode JSON from Request Body to User Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Database Query
	_, err := helpers.DB.Exec("INSERT INTO users (name, email) VALUE (?, ?)", user.Name, user.Email)
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
		// Get ID Parameters From URI Then Convert it to Integer
		userID, err := strconv.Atoi(params["id"])
		if err == nil {
			var user models.User
			var response routers.Response

			// Decode JSON from Request Body to User Data
			// Use _ As Temporary Variable
			_ = json.NewDecoder(r.Body).Decode(&user)

			// Database Query
			_, err := helpers.DB.Exec("UPDATE users SET name=?, email=? WHERE id=? LIMIT 1", user.Name, user.Email, userID)
			if err != nil {
				log.Print(err)
			}

			// Set Response Data
			response.Status = true
			response.Code = http.StatusCreated
			response.Message = "Success"

			// Write Response Data to HTTP
			routers.ResponseWrite(w, response.Code, response)
		} else {
			routers.ResponseInternalError(w)
		}
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
		// Get ID Parameters From URI Then Convert it to Integer
		userID, err := strconv.Atoi(params["id"])
		if err == nil {
			var response routers.Response

			// Database Query
			_, err := helpers.DB.Query("DELETE FROM users WHERE id=? LIMIT 1", userID)
			if err != nil {
				log.Print(err)
			}

			// Set Response Data
			response.Status = true
			response.Code = http.StatusOK
			response.Message = "Success"

			// Write Response Data to HTTP
			routers.ResponseWrite(w, response.Code, response)
		} else {
			routers.ResponseInternalError(w)
		}
	}
}
