package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//arrays declaration
	var a = [5]int{1, 3, 2, 5, 4}
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[3])

	sort.Ints(a[:]) //Ints is used for integers sorting
	fmt.Println("After sorting :", a)

	//without giving array length we also prints arrays using  ... dots
	var arr = [...]string{"new", "kavya"}
	fmt.Println(arr)
	arr[1] = "mouni"
	fmt.Println(arr)
	fmt.Println(len(arr))
	var i int
	for i = 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	sort.Strings(arr[:]) //Strings is used for integers sorting
	fmt.Println("After sorting :", arr)
	fmt.Println("Absolute :", math.Abs(-10))
	fmt.Println("max value :", math.Max(20, 30))
	fmt.Println("min value :", math.Min(50, 10))
	fmt.Println("mod value :", math.Mod(10, 2))
	fmt.Println("power function :", math.Pow(10, 2))
	fmt.Println("sqrt function :", math.Sqrt(64))
	fmt.Println("cbrt function :", math.Cbrt(64))
	fmt.Println("floor function :", math.Floor(4.7))
	fmt.Println("round function :", math.Round(4.7))

	//appending values to arrays
	var str = []string{}
	str = append(str, "new", "mouni", "kavya", "karuna", "sai", "ram")
	fmt.Println(str)
	fmt.Println(str[1:3])
	fmt.Println(str[3:])

	//range function .it returns 2 values index and value
	for index, value := range str {
		fmt.Println(index, value)
	}

	/*make functions The make function is a built-in function in Go programming language used to initialize special data types.
	What make does
	It creates and initializes:
		slices
		maps
		channels
		It does not work for normal variables like int, string, etc.
		make(type, length, capacity)*/
	var s = make([]int, 5, 10)
	fmt.Println(s, cap(s))

	/*A map is just:
	key → value storage (like a dictionary)*/
	m := make(map[string]int)

	m["age"] = 21
	m["marks"] = 90

	fmt.Println(m)
	m["age"] = 30
	fmt.Println(m)
	delete(m, "age")
	fmt.Println(m)
	fmt.Println(m["age"])
}
