package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//creating web server
/*func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Web server created success")
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server is not running")
	}
}
*/

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/add", Addfunc)
	http.HandleFunc("/api", handleMethod)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server is not running")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server created ")
}

// passing query prameters
func Addfunc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	a := r.FormValue("a")
	b := r.FormValue("b")

	a1, err1 := strconv.Atoi(a)
	b1, err2 := strconv.Atoi(b)

	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "adition of %d and %d is %d", a1, b1, a1+b1)
}

// Methods
func handleMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "Get method")
	case http.MethodPost:
		fmt.Fprint(w, "Post method")
	case http.MethodPut:
		fmt.Fprint(w, "Put method")
	case http.MethodDelete:
		fmt.Fprint(w, "Delete method")
	}
}
