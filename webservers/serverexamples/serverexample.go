package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name          string
	Qualification string
	Age           int
}

var person []Person

func AddPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "post method is only allowed", http.StatusMethodNotAllowed)
		return
	}

	var p Person
	json.NewDecoder(r.Body).Decode(&p)
	person = append(person, p)
	json.NewEncoder(w).Encode(person)
}

func getPersons(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "get method only allowed", http.StatusMethodNotAllowed)
		return
	}

	json.NewEncoder(w).Encode(person)
}

func getPersonByName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Get Method only allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	for _, p := range person {
		if p.Name == name {
			json.NewEncoder(w).Encode(&p)
			return
		}
	}
	http.Error(w, "person not found", http.StatusNotAcceptable)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "only delete method is allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	for i, p := range person {
		if p.Name == name {
			person = append(person[:i], person[i+1:]...) //means you are deleting the element at index i by taking all elements before it and after it, and joining them together into a new slice (list), so the item at index i is removed.
			fmt.Fprint(w, "Person deleted successfully")
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/addPersons", AddPerson)
	http.HandleFunc("/persons", getPersons)
	http.HandleFunc("/getPerosnName", getPersonByName)
	http.HandleFunc("/deletePerson", deletePerson)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server not running")
	}
}
