package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// User struct
type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

// Users var as a slice of user struct
var Users []User

func main() {
	// Init router
	router := mux.NewRouter()

	// Route Handlers
	router.HandleFunc("/api/user", getUsers).Methods("GET")
	router.HandleFunc("/api/user/{Username}", getUser).Methods("GET")
	router.HandleFunc("/api/user", createUser).Methods("POST")
	router.HandleFunc("/api/user/{Username}", updateUser).Methods("PUT")
	router.HandleFunc("/api/user/{Username}", deleteUser).Methods("DELETE")

	// Start Server
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r) // get parameters
	// find username
	for _, item := range Users {
		if item.Username == param["Username"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(User{})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	Users = append(Users, newUser)
	json.NewEncoder(w).Encode(Users)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r) // get parameters
	for i, item := range Users {
		if item.Username == param["Username"] {
			Users = append(Users[:i], Users[i+1:]...)
			break
		}
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r) // get parameters
	for i, item := range Users {
		if item.Username == param["Username"] {
			Users = append(Users[:i], Users[i+1:]...)
			var newUser User
			json.NewDecoder(r.Body).Decode(&newUser)
			newUser.Username = param["Username"]
			Users = append(Users, newUser)
			json.NewEncoder(w).Encode(Users)
			return
		}
	}

}
