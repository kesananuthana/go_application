package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

type Products struct {
	Pid   string
	Name  string
	Price int
}

var conn *pgx.Conn

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

func products(w http.ResponseWriter, r *http.Request) {
	row, err := conn.Query(
		context.Background(),
		"select * from products",
	)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	var product []Products
	for row.Next() {
		var p Products
		err := row.Scan(&p.Pid, &p.Name, &p.Price)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		product = append(product, p)
	}
	json.NewEncoder(w).Encode(product)
}

func productId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pid := params["pid"]

	row := conn.QueryRow(
		context.Background(),
		"select * from products where pid = $1",
		pid,
	)

	var p Products
	err := row.Scan(&p.Pid, &p.Name, &p.Price)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Product not found",
		})
		return
	}
	json.NewEncoder(w).Encode(p)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pid := params["pid"]

	var p Products
	json.NewDecoder(r.Body).Decode(&p)

	row, err := conn.Exec(
		context.Background(),
		"update products set name =$2, price =$3 where pid =$1",
		pid,
		p.Name,
		p.Price,
	)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	if row.RowsAffected() == 0 {
		fmt.Fprint(w, "Product not found")
		return
	}

	fmt.Fprint(w, "Product updated successfully")
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pid := params["pid"]

	row, err := conn.Exec(
		context.Background(),
		"delete from products where pid = $1",
		pid,
	)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if row.RowsAffected() == 0 {
		fmt.Fprint(w, "Product not found")
		return
	}
	fmt.Fprint(w, "Pproduct deleted successfully")
}

func deleteAll(w http.ResponseWriter, r *http.Request) {
	row, err := conn.Exec(
		context.Background(),
		"delete from products",
	)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if row.RowsAffected() == 0 {
		fmt.Fprint(w, "No Products found")
		return
	}

	fmt.Fprint(w, "All Prdocuts deleted successfully")
}

func main() {
	connStr := "postgres://postgres:nuthana@localhost:5432/users"
	var err error
	conn, err = pgx.Connect(context.Background(), connStr)

	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println("Database connected successfully", conn)

	router := mux.NewRouter()
	router.HandleFunc("/products", withCORS(products)).Methods("GET")
	router.HandleFunc("/productId/{pid}", withCORS(productId)).Methods("GET")
	router.HandleFunc("/updateProduct/{pid}", updateProduct).Methods("PUT")
	router.HandleFunc("/deleteProduct/{pid}", deleteProduct).Methods("DELETE")
	router.HandleFunc("/deleteAllProducts", deleteAll).Methods("DELETE")
	http.ListenAndServe(":8080", router)
	fmt.Print("Server running on port 8080")
}
