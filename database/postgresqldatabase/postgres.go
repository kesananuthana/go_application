package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

type UserDetails struct {
	Id    string
	Email string
	Name  string
}

func main() {
	//connStr := "postgres://postgres:nuthana@localhost:5432/users"

	connStr := os.Getenv("db_url")

	var err error

	conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Print("connection failed", err)
	}
	defer conn.Close(context.Background())
	fmt.Print("database connected successfully", conn)

	router := mux.NewRouter()

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/addUser", addUsers).Methods("POSt")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/usersById/{id}", getUserById).Methods("GET")
	router.HandleFunc("/updateUser/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/deleteUsers", deleteUsers).Methods("DELETE")
	router.HandleFunc("/deleteUserById/{id}", deleteUserById).Methods("DELETE")
	http.ListenAndServe(":8080", router)
	fmt.Print("Server running on port 8080")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func addUsers(w http.ResponseWriter, r *http.Request) {
	var user UserDetails
	json.NewDecoder(r.Body).Decode(&user)
	_, err := conn.Exec(
		context.Background(),
		"INSERT INTO userdetails (id, email, name) VALUES ($1, $2, $3)",
		user.Id,
		user.Email,
		user.Name,
	)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, "user added successfully")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := conn.Query(
		context.Background(),
		"Select * from userdetails",
	)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	var userdetails []UserDetails

	for rows.Next() {

		var u UserDetails

		err := rows.Scan(&u.Id, &u.Email, &u.Name) //Scan() copies database column values into Go struct fields.

		if err != nil {
			fmt.Println(err)
			return
		}

		userdetails = append(userdetails, u)
	}

	json.NewEncoder(w).Encode(userdetails)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	row := conn.QueryRow(
		context.Background(),
		"Select * from userdetails where id = $1",
		id,
	)

	var u UserDetails
	err := row.Scan(&u.Id, &u.Email, &u.Name)

	if err != nil {
		fmt.Fprint(w, "Student not found")
		return
	}

	json.NewEncoder(w).Encode(u)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var user UserDetails
	json.NewDecoder(r.Body).Decode(&user)
	res, err := conn.Exec(
		context.Background(),
		"update userdetails set email = $2 , name = $3 where id = $1",
		id,
		user.Email,
		user.Name,
	)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if res.RowsAffected() == 0 {
		fmt.Fprint(w, "User not exist")
		return
	}
	fmt.Fprint(w, "user updated successfully")
}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	_, err := conn.Exec(
		context.Background(),
		"Delete from userdetails",
	)

	if err != nil {
		fmt.Fprint(w, "No Users exist")
		return
	}
	fmt.Fprint(w, "All users deleted successfully")
}

func deleteUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	res, err := conn.Exec(
		context.Background(),
		"delete from userdetails where id = $1",
		id,
	)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	if res.RowsAffected() == 0 {
		fmt.Fprint(w, "User not exist")
		return
	}

	fmt.Fprint(w, "user deleted successfully")
}
