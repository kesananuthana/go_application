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

type Products struct {
	Pid   string
	Name  string
	Price int
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func addProducts(w http.ResponseWriter, r *http.Request) {
	var products Products
	json.NewDecoder(r.Body).Decode(&products)
	_, err := conn.Exec(
		context.Background(),
		"Insert into products(pid,name,price) values($1,$2,$3)",
		products.Pid,
		products.Name,
		products.Price,
	)
	if err != nil {
		fmt.Fprint(w, err)
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product added successfully",
	})
}
func getproducts(w http.ResponseWriter, r *http.Request) {
	rows, err := conn.Query(
		context.Background(),
		"select * from products",
	)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	var products []Products
	for rows.Next() {
		var p Products
		err := rows.Scan(&p.Pid, &p.Name, &p.Price)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		products = append(products, p)
	}
	json.NewEncoder(w).Encode(products)
}

func productsByid(w http.ResponseWriter, r *http.Request) {
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

func productsByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	fmt.Println("Received name:", name)

	row := conn.QueryRow(
		context.Background(),
		"select * from products where name = $1",
		name,
	)
	var p Products
	err := row.Scan(&p.Pid, &p.Name, &p.Price)

	if err != nil {

		fmt.Println("Scan error:", err)
		fmt.Fprint(w, "product not found")
		return
	}
	json.NewEncoder(w).Encode(p)
}

func updateProducts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pid := params["pid"]

	var p Products
	json.NewDecoder(r.Body).Decode(&p)
	res, err := conn.Exec(
		context.Background(),
		"update products set name = $2, price =$3 where pid =$1",
		pid,
		p.Name,
		p.Price,
	)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if res.RowsAffected() == 0 {
		fmt.Fprint(w, "Product not found")
		return
	}
	fmt.Fprint(w, "Product updated successfully")
}

func updateProductById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pid := params["pid"]

	var p Products
	json.NewDecoder(r.Body).Decode(&p)
	res, err := conn.Exec(
		context.Background(),
		"update products set name = $2, price = $3 where pid = $1",
		pid,
		p.Name,
		p.Price,
	)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if res.RowsAffected() == 0 {
		fmt.Fprint(w, "Product not found")
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product updated successfully",
	})
}

func deleteProductsById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pid := params["pid"]

	res, err := conn.Exec(
		context.Background(),
		"delete from products where pid =$1",
		pid,
	)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if res.RowsAffected() == 0 {
		fmt.Fprint(w, "Product not exist")
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product deleted successfully",
	})
}
func deleteProducts(w http.ResponseWriter, r *http.Request) {
	_, err := conn.Exec(
		context.Background(),
		"delete from products",
	)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "All Products deleted",
	})
}

func main() {
	//connStr := "postgres://postgres:nuthana@localhost:5432/users"
	connStr := os.Getenv("db_url")

	var err error
	conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Print("Connection failed", err)
	}
	fmt.Print("Database connected")
	router := mux.NewRouter()
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})
	router.HandleFunc("/", greet).Methods("GET")
	router.HandleFunc("/addProducts", addProducts).Methods("POST", "OPTIONS")
	router.HandleFunc("/products", getproducts).Methods("GET")
	router.HandleFunc("/ProductById/{pid}", productsByid).Methods("GET")
	router.HandleFunc("/ProductByName/{name}", productsByName).Methods("GET")
	router.HandleFunc("/updateProducts/{pid}", updateProducts).Methods("PUT")
	router.HandleFunc("/updateProductById/{pid}", updateProductById).Methods("PUT", "OPTIONS")
	router.HandleFunc("/deletProducts/{pid}", deleteProductsById).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/deleteProducts", deleteProducts).Methods("DELETE", "OPTIONS")
	http.ListenAndServe(":8080", router)
	fmt.Print("Server running on port 8080")
}

/*corsHandler := handlers.CORS(
	handlers.AllowedOrigins([]string{"*"}),
	handlers.AllowedMethods([]string{
		"GET", "POST", "PUT", "DELETE", "OPTIONS",
	}),
	handlers.AllowedHeaders([]string{
		"Content-Type", "Authorization",
	}),
)*/
