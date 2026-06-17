package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	Id   string
	Name string
}

var students = []Student{
	{Id: "1", Name: "Ganesh"},
	{Id: "2", Name: "Rahul"},
	{Id: "3", Name: "Mouni"},
	{Id: "4", Name: "New"},
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Go routers")
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(students)
}

func getStudentById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) //mux.Vars(r) is used to get dynamic values (parameters) from the URL in Go.
	id := params["id"]
	for _, s := range students {
		if s.Id == id {
			json.NewEncoder(w).Encode(s)
			return
		}
	}
	fmt.Fprint(w, "No student found")
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var update Student
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid request body")
		return
	}
	for i, s := range students {
		if s.Id == id {
			students[i].Name = update.Name
			json.NewEncoder(w).Encode(students[i])
			return
		}
	}
	fmt.Fprint(w, "No student found")
}

func deleteStudentById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for i, s := range students {
		if s.Id == id {
			students = append(students[:i], students[i+1:]...)
			fmt.Fprint(w, "Student deleted successfully")
			return
		}
	}
	fmt.Fprint(w, "Student not found")
}

func deleteStudents(w http.ResponseWriter, r *http.Request) {

	students = []Student{}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Students deleted successfully",
	})
}

func main() {
	router := mux.NewRouter() //Creates router object.

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/students", getStudents).Methods("GET")
	router.HandleFunc("/studentById/{id}", getStudentById).Methods("GET")
	router.HandleFunc("/updateStudent/{id}", updateStudent).Methods("PUT")
	router.HandleFunc("/deleteStudent/{id}", deleteStudentById).Methods("DELETE")
	router.HandleFunc("/deleteStudent", deleteStudents).Methods("DELETE")
	fmt.Print("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
