package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"
)

func hello(a, b int) (int, int) {
	c := a + b
	d := a * b
	fmt.Println("\n hii")
	return c, d
}

type Student struct {
	Name  string
	Age   int
	Class string
}

var student []Student

func addStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post method only allowed", http.StatusMethodNotAllowed)
		return
	}
	var s Student
	json.NewDecoder(r.Body).Decode(&s)
	student = append(student, s)
	json.NewEncoder(w).Encode(student)
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "get method only accepted", http.StatusMethodNotAllowed)
	}

	message := make(map[string]string)
	message["message"] = "No Strudents exist"
	if len(student) == 0 {
		json.NewEncoder(w).Encode(message)
	}

	json.NewEncoder(w).Encode(student)
}

func getByStudentName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "post method only accepted", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	for _, s := range student {
		if s.Name == name {
			json.NewEncoder(w).Encode(s)
			return
		}
	}

	message := make(map[string]string)
	message["message"] = "Student not found"
	json.NewEncoder(w).Encode(message)
}

func deleteStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Delete method only accepted", http.StatusMethodNotAllowed)
		return
	}
	student = []Student{}
	message := make(map[string]string)
	message["message"] = "All Students deleted"
	json.NewEncoder(w).Encode(message)
}

func deleteStudentByName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, " delete method only accepted", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	message := make(map[string]string)
	message["message"] = "Student deleted"
	for i, s := range student {
		if s.Name == name {
			student = append(student[:i], student[i+1:]...)
			json.NewEncoder(w).Encode(message)
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

func updatestudentByName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "put method only accepted", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	var update Student
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, "Invalid Data", http.StatusNotFound)
	}

	for i, s := range student {
		if s.Name == name {
			student[i].Name = update.Name
			student[i].Age = update.Age
			student[i].Class = update.Class

			json.NewEncoder(w).Encode(student[i])
			return
		}
	}
	http.Error(w, "student not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/addStudents", addStudents)
	http.HandleFunc("/students", getStudents)
	http.HandleFunc("/studentName", getByStudentName)
	http.HandleFunc("/deleteStudents", deleteStudents)
	http.HandleFunc("/deleteStudent", deleteStudentByName)
	http.HandleFunc("/updateStudent", updatestudentByName)
	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		fmt.Println("Server not running")
	}

	var flag bool
	n := 10
	n1 := 20
	for i := n; i <= n1; i++ {
		flag = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}

	var a = []int{5, 2, 11, 3, 4, 9}
	max := a[0]
	min := a[0]
	for i := 0; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		} else if min > a[i] {
			min = a[i]
		}
	}
	a = append(a, 8)
	fmt.Println("max element is ", max)
	fmt.Println("min element is", min)
	for i := range a {
		fmt.Println(a[i])
	}
	s := append(a, 5)
	fmt.Println(s)
	fmt.Println(a[2:4])
	fmt.Println(a[:2])

	m := make(map[string]string)
	m["name"] = "mouni"
	m["age"] = "20"
	fmt.Println(m)

	//short ways to create map
	m1 := map[string]string{
		"name": "mouni",
		"age":  "20",
	}
	fmt.Println(m1)

	for key, value := range m {
		fmt.Println(key, value)
	}

	fmt.Println(time.Now(), time.Date(2024, time.January, 1, 0, 0, 0, 0, time.Local))
	data := "new,mouni,kavya"
	fmt.Println(data)
	data = strings.ToUpper(data)
	fmt.Println(data)
	parts := strings.Split(data, ",")
	fmt.Println(parts)
	name := strings.Replace(data, "NEW", "likki", -1)
	fmt.Println(name)
	s1 := "hello"
	s2 := "world"
	fmt.Println(s1 + " " + s2) //strign concatenation
	fmt.Println(s1 == s2)
	fmt.Println(s1 < s2)
	fmt.Println(strings.Compare(s1, s2))
	fmt.Println(strings.Contains(s1, "ll"))

	d := "mouni"
	for i := len(d) - 1; i >= 0; i-- {
		fmt.Printf("%c", d[i])
	}
	var res, res1 = hello(10, 20)
	fmt.Println(res, " ", res1)

	var student = Student{"kavya", 20, "6th class"}
	fmt.Println("student details", student)

	data1 := []string{"hello", "world", "hii"}
	fmt.Println(data1)
	maxval := slices.Max(data1)
	minVal := slices.Min(data1)
	fmt.Println("maxval is", maxval)
	fmt.Println("minVal is", minVal)
	data1 = append(data1, "apple")
	for _, i := range data1 {
		fmt.Println(i)
	}

	res, err := div(10, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	var i, j int
	var k string
	fmt.Println("Enter i,j ,k values,")
	fmt.Scan(&i, &j, &k)
	fmt.Println(cal(i, j, k))
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Not divisible")
	}
	return a / b, nil
}

func cal(i, j int, k string) int {
	var c int
	switch k {
	case "+":
		c = i + j
	case "-":
		c = i - j
	case "*":
		c = i * j
	default:
		c = 0
	}
	return c
}
