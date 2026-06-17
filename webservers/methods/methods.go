package main

import (
	"encoding/json" // it is used for converting go to json
	"fmt"
	"net/http"
)

type Student struct {
	Name string
	Age  int
}

var students []Student //it stores students records

// post method
func addStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	var s Student                      //it stores incoming json data
	json.NewDecoder(r.Body).Decode(&s) // it takes json data and Converts into Go struct s

	students = append(students, s)

	json.NewEncoder(w).Encode(s) //Sends back the inserted student as JSON
}

// get method
func getStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only get allowed", http.StatusMethodNotAllowed)
		return
	}
	json.NewEncoder(w).Encode(students)
}

// get data by Name
func getByStudentName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only get method is allowed", http.StatusMethodNotAllowed)
	}

	name := r.URL.Query().Get("name") //Read query parameter from URL

	for _, s := range students {
		if s.Name == name {
			json.NewEncoder(w).Encode(s)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusMethodNotAllowed)
}

// update data
func updateStudentData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "put method oly allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	var update Student
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	for i, s := range students {
		if s.Name == name {
			students[i].Name = update.Name
			students[i].Age = update.Age

			json.NewEncoder(w).Encode(students[i])
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}

// delete method
func deleteStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "only delete method is allowed", http.StatusMethodNotAllowed)
	}

	students = []Student{}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "All students deleted successfully",
	})
}

func main() {

	http.HandleFunc("/addStudents", addStudents)
	http.HandleFunc("/students", getStudents)
	http.HandleFunc("/studentName", getByStudentName)
	http.HandleFunc("/updateStudent", updateStudentData)
	http.HandleFunc("/deleteStudents", deleteStudents)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server not running")
	}
}

//in the above w passing in all functions.beacuse Go, send response back to browser
//in above code it use NewEncoder and Encode.
/*NewEncoder(w) → create encoder This creates an encoder object
It prepares something that can write JSON
It needs a destination (w, file, etc.)
.Encode(students) → send JSON response (This converts Go data → JSON and writes it)*/
