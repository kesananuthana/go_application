package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"         //MongoDB driver
	"go.mongodb.org/mongo-driver/mongo/options" //mongodb connection settings
)

type Student struct {
	Id   string
	Name string
}

var collection *mongo.Collection //This variable stores the MongoDB collection.

func main() {

	// connect mongodb
	client, err := mongo.Connect(
		context.TODO(), //Context controls: timeout, cancellation, request ,lifecycle
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("MongoDB Connected")

	// database and collection
	collection = client.Database("school").Collection("students")

	// route
	http.HandleFunc("/students", addStudent)
	http.HandleFunc("/studentsData", getStudents)
	http.HandleFunc("/updateStudent", updateStudent)
	http.HandleFunc("/deleteStudent", deleteStudent)
	http.HandleFunc("/deleteAll", deleteAllStudents)
	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)
}

// POST API
func addStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post method only accepted", http.StatusMethodNotAllowed)
	}

	var student Student

	// json -> struct
	json.NewDecoder(r.Body).Decode(&student)

	// insert into mongodb
	_, err := collection.InsertOne(context.TODO(), student)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprint(w, "student added successfully")
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "get method only accepted", http.StatusMethodNotAllowed)
		return
	}
	cursor, err := collection.Find(context.TODO(), bson.M{}) //bson.M{} means: “Give me ALL documents in collection”
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer cursor.Close(context.TODO()) //close cursor AFTER function finishes

	var students []Student

	for cursor.Next(context.TODO()) { //Reads documents one by one from MongoDB result set
		var student Student     //Temporary variable to hold one document at a time
		cursor.Decode(&student) // Converts MongoDB BSON document into Go struct
		students = append(students, student)
	}

	json.NewEncoder(w).Encode(students)
	if len(students) == 0 {
		fmt.Fprint(w, "No Students exist")
	}
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "PUT only allowed", http.StatusMethodNotAllowed)
		return
	}

	var student Student
	json.NewDecoder(r.Body).Decode(&student)

	filter := bson.M{"id": student.Id}

	update := bson.M{
		"$set": bson.M{
			"name": student.Name,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "DELETE only allowed", http.StatusMethodNotAllowed)
		return
	}

	var student Student
	json.NewDecoder(r.Body).Decode(&student)

	filter := bson.M{"id": student.Id}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func deleteAllStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Delete method only accepted", http.StatusMethodNotAllowed)
		return
	}

	res, err := collection.DeleteMany(
		context.TODO(),
		bson.M{},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintf(
		w,
		"%d students deleted",
		res.DeletedCount,
	)
}

/*bson.M{} It is used for:
filters
queries
updates
conditions in MongoDB.
bson.M{} is used to create MongoDB query filters and conditions in Go
*/
