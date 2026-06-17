package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

type Person struct {
	Id   string
	Name string
	Age  int
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	json.NewDecoder(r.Body).Decode(&person)

	_, err := collection.InsertOne(context.TODO(), person)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	fmt.Fprint(w, "Student added successfully")
}

func Persons(w http.ResponseWriter, r *http.Request) {
	data, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	defer data.Close(context.TODO())
	var person []Person
	for data.Next(context.TODO()) {
		var p Person
		data.Decode(&p)
		person = append(person, p)
	}

	json.NewEncoder(w).Encode(person)
	if len(person) == 0 {
		fmt.Fprint(w, "No students exist")
	}
}

func getPersonsById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	filter := bson.M{"id": id}
	var res Person
	err := collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		fmt.Fprint(w, "student not found")
		return
	}
	json.NewEncoder(w).Encode(res)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	var p Person
	json.NewDecoder(r.Body).Decode(&p)
	params := mux.Vars(r)
	id := params["id"]
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"name": p.Name,
			"age":  p.Age,
		},
	}

	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	if res.MatchedCount == 0 {
		fmt.Fprint(w, "student not found")
		return
	}
	json.NewEncoder(w).Encode(res)
}

func deletePersons(w http.ResponseWriter, r *http.Request) {
	data, err := collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	fmt.Fprintf(
		w, "%d students deleted successfully", data.DeletedCount,
	)
}

func deletePersonById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	filter := bson.M{"id": id}
	data, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
	fmt.Fprint(w, "No students found")
}

func main() {

	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)

	if err != nil {
		fmt.Println("Mongodb not connected")
	}
	fmt.Println("mongoDb connected")
	collection = client.Database("school").Collection("person")

	router := mux.NewRouter()
	router.HandleFunc("/", addPerson).Methods("POST")
	router.HandleFunc("/persons", Persons).Methods("GET")
	router.HandleFunc("/personById/{id}", getPersonsById).Methods("GET")
	router.HandleFunc("/updatePerson/{id}", updatePerson).Methods("PUT")
	router.HandleFunc("/deletePersons", deletePersons).Methods("DELETE")
	router.HandleFunc("/deletePersonByID/{id}", deletePersonById).Methods("DELETE")
	http.ListenAndServe(":8080", router)

	fmt.Print("Server running on port 8080")
}
