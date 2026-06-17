package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Users struct {
	Email    string
	Password string
}

var users []Users

// Creates a middleware function.
// Middleware means: run some code before actual API function.
func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*") //Allows frontend from any origin to access backend.
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type") //Allows frontend to send Content-Type header.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		} //Browser first sends OPTIONS request for checking permissions. If request is OPTIONS: send status 200 OK stop function using return.

		next(w, r) //Calls actual API function after CORS handling.
	}
}

func addLoginDetails(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "only post method allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var u Users
	json.NewDecoder(r.Body).Decode(&u)
	users = append(users, u)
	json.NewEncoder(w).Encode(users)
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "only get method allowed", http.StatusMethodNotAllowed)
	}

	json.NewEncoder(w).Encode(users)
}

func getByEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "get method only allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.URL.Query().Get("email")
	for _, u := range users {
		if u.Email == email {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func updateUserDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "put method only allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.URL.Query().Get("email")
	var update Users
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	for i, u := range users {
		if u.Email == email {
			users[i].Email = update.Email
			users[i].Password = update.Password
			json.NewEncoder(w).Encode(users)
			return
		}
	}

	http.Error(w, "Usernot found", http.StatusNotFound)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, " delete method only accepted", http.StatusMethodNotAllowed)
		return
	}

	email := r.URL.Query().Get("email")
	for i, u := range users {
		if u.Email == email {
			users = append(users[:i], users[i+1:]...)
			fmt.Print(w, "user deleted successfully")
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/addUsers", withCORS(addLoginDetails))
	http.HandleFunc("/users", withCORS(getUsers))
	http.HandleFunc("/useremail", withCORS(getByEmail))
	http.HandleFunc("/updateUser", withCORS(updateUserDetails))
	http.HandleFunc("/removeUser", withCORS(deleteUser))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server not running")
	}
}
