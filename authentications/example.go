package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

type User struct {
	Email    string
	Password string
}

var user []User
var jwtKey = []byte("my_secret_key") //it creating token verifying token

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*") //Allows frontend from any origin to access backend.
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST,PUT,DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,  Authorization") //Allows frontend to send Content-Type header.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		} //Browser first sends OPTIONS request for checking permissions. If request is OPTIONS: send status 200 OK stop function using return.

		next(w, r) //Calls actual API function after CORS handling.
	}
}

// authentication middleware
func authMiddleware(next http.HandlerFunc) http.HandlerFunc { //next means → next API function

	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization") // Reads token from request header.

		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenString = tokenString[len("Bearer "):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { // This verifies token.
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r) // If token valid: allow request to continue .This runs actual API function.
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	for _, newUser := range user {
		if newUser.Email == u.Email {
			json.NewEncoder(w).Encode(map[string]string{
				"message": "User already exist",
			})
			return
		}
	}
	user = append(user, u)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

func getRegisterUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(user)
}

func getLoginUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(user)
}

func generateToken(email string) (string, error) {

	claims := jwt.MapClaims{ // Claims = data stored inside JWT.
		//Your token stores : email & expiry time
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // 1 day expiry
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //Creates token object. used HS256 algorithm

	return token.SignedString(jwtKey) //Signs token using secret key.
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser User
	json.NewDecoder(r.Body).Decode(&loginUser)
	for _, user := range user {
		if user.Email == loginUser.Email && user.Password == loginUser.Password {
			token, err := generateToken(loginUser.Email)
			if err != nil {
				http.Error(w, "Error generating token", http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(map[string]string{
				"token": token,
			})
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Invalid email or password",
	})
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]
	decodedEmail, _ := url.QueryUnescape(email)

	var u User
	json.NewDecoder(r.Body).Decode(&u)
	for i, u1 := range user {
		if u1.Email == decodedEmail {
			user[i].Email = u.Email
			user[i].Password = u.Password
			json.NewEncoder(w).Encode(user[i])
			return
		}
	}
	fmt.Fprint(w, "User not found")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]
	decodedEmail, _ := url.QueryUnescape(email)
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	for i, s := range user {
		if s.Email == decodedEmail {
			user = append(user[:i], user[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "User deleted successfully",
			})
			return
		}
	}
	fmt.Fprint(w, "User not found")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/registerUser", withCORS(createUser)).Methods("POST", "OPTIONS")
	router.HandleFunc("/getRegisterUsers", withCORS(authMiddleware(getRegisterUsers))).Methods("GET", "OPTIONS")
	router.HandleFunc("/loginUser", withCORS(loginUser)).Methods("POST", "OPTIONS")
	router.HandleFunc("/getloginUsers", withCORS(authMiddleware(getLoginUsers))).Methods("GET", "OPTIONS")
	router.HandleFunc("/updateUser/{email}", withCORS(updateUser)).Methods("PUT", "OPTIONS")
	router.HandleFunc("/deleteUser/{email}", withCORS(deleteUser)).Methods("DELETE", "OPTIONS")
	http.ListenAndServe(":8080", router)
	fmt.Print("Server running on port 8080")
}
