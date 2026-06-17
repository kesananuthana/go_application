package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Id      string
	Name    string
	Address Address //inside Person, the Address field must store an Address object.
}

type Address struct {
	City    string
	Pincode int
}

var person []Person

func createPerson(w http.ResponseWriter, r *http.Request) {
	var p Person
	json.NewDecoder(r.Body).Decode(&p)

	person = append(person, p)
	json.NewEncoder(w).Encode(person)

}

func Persons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(person)
}

func personById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for _, p := range person {
		if p.Id == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	fmt.Fprint(w, "Person not found")
}

func getAddress(w http.ResponseWriter, r *http.Request) {
	var address = []Address{}
	for _, a := range person {
		address = append(address, a.Address)
		json.NewEncoder(w).Encode(address)
	}
}

func getAddressById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var address = []Address{}

	for _, p := range person {
		if p.Id == id {
			address = append(address, p.Address)
			json.NewEncoder(w).Encode(address)
		}
	}
}

func updatePerson(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	var persons Person
	json.NewDecoder(r.Body).Decode(&persons)
	for i, p := range person {
		if p.Id == id {
			person[i].Name = persons.Name
			person[i].Address.City = persons.Address.City
			person[i].Address.Pincode = persons.Address.Pincode
			json.NewEncoder(w).Encode(person[i])
			return
		}
	}
	fmt.Fprint(w, "Person not found")
}

func updateAddress(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var persons Person
	json.NewDecoder(r.Body).Decode(&persons)
	for i, p := range person {
		if p.Id == id {
			person[i].Address.City = persons.Address.City
			person[i].Address.Pincode = persons.Address.Pincode
			json.NewEncoder(w).Encode(person[i])
			return
		}
	}
}

func deletePersonById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for i, p := range person {
		if p.Id == id {
			person = append(person[:i], person[i+1:]...)
			fmt.Fprint(w, "Person deleted successfully")
			return
		}
	}
	fmt.Fprint(w, "Person not found")
}

func deleteAll(w http.ResponseWriter, r *http.Request) {
	person = []Person{}
	fmt.Fprint(w, "Persons deleted successfully")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", createPerson).Methods("POST")
	router.HandleFunc("/persons", Persons).Methods("GET")
	router.HandleFunc("/PersonById/{id}", personById).Methods("GET")
	router.HandleFunc("/address", getAddress).Methods("GET")
	router.HandleFunc("/addressById/{id}", getAddressById).Methods("GET")
	router.HandleFunc("/updateperson/{id}", updatePerson).Methods("PUT")
	router.HandleFunc("/updateAddress/{id}", updateAddress).Methods("PUT")
	router.HandleFunc("/deletePersonId/{id}", deletePersonById).Methods("DELETE")
	router.HandleFunc("/deleteAll", deleteAll).Methods("DELETE")
	http.ListenAndServe(":8080", router)
	fmt.Print("Server running on port 8080")
}
