package main

import "fmt"

//this main function
//taking inpputs from user
func main() {
	wish()
	wish()
	add(10, 20, "new")
	var c, d, e = sub(40, 30)
	defer fmt.Print(c, d, e)
	var x = 52.5
	fmt.Print("\n", int(x))
	var name string
	var city string
	var a int
	fmt.Print("Enter the name")
	//for reading inpputs use scaln and &
	fmt.Scanln(&name, &city, &a)
	fmt.Println("my name is ", name)
	fmt.Println(city)
	fmt.Println(a)
	switch name {
	case "new":
		fmt.Println("new")
	case "abc":
		fmt.Println("abc")
	default:
		fmt.Println("default value")
	}

	fmt.Println("sum", sum(1, 2, 3, 4))
	fmt.Println("sum is ", sum(1, 2, 3))
}

func wish() {
	fmt.Println("Welcome")
}

//functions with parameters
func add(a int, b int, name string) {
	a = 20
	var c = a + b
	var d float64 = float64(c)
	fmt.Print("value of d", d)
	fmt.Println("\n hello", name)
}

//functions with return values
func sub(a int, b int) (int, int, int) {
	var c = a - b
	var d = a * b
	var e = a / b
	return c, d, e
}

//varidic functions : Variadic functions in Go are functions that accept a variable number of arguments of the same type.
func sum(nums ...int) int {
	total := 0
	var i int
	for i = 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}
