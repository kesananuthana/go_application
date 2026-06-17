package main

import (
	"fmt"
	"learn/packages" //importing packages
	"learn/packages/util"
	"strconv"
)

func main() {
	var c int
	var e float64
	var class string
	var a int = 10
	var b float64 = 50.7
	var name string = "new"
	var istrue bool = false
	fmt.Print("value of a is : ", a)
	fmt.Println("\n b value is", b)
	fmt.Print("\n name is ,", name)
	fmt.Print("\n bool :", istrue)
	fmt.Print(c)
	fmt.Print(e)
	fmt.Print(class)
	a = 20
	fmt.Println(a)
	age := 25
	print(age)
	var x, y, z int = 10, 20, 30
	fmt.Print(x, y, z)
	//for chceking datatype
	fmt.Printf("\n %T", name)
	fmt.Printf(" \t %v", x)
	const pi = 3.14
	fmt.Print("\n", pi)
	//for multiple constants declaration
	const (
		d    = 23
		city = "hyd"
	)
	fmt.Println("\n", d, city)

	//integer to string conversition
	var i int = 456
	var s string = strconv.Itoa(i) // Itoa is used for converting int to alphabets
	fmt.Print(s)

	//string to integer conversition
	var str string = "123"
	var n, err = strconv.Atoi(str) //Atoi is used for converting alphabets to int
	if err == nil {
		fmt.Println("\n", n)
	} else {
		fmt.Println(err)
	}
	fmt.Printf("%T", n)
	fmt.Println(a == age)
	for i = 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Print(i)
	}

	b = 5
	for b >= 1 {
		fmt.Println("value of b", b)
		b--
	}
	n = 8
	c = 0
	for i = 1; i <= n; i++ {
		if n%i == 0 {
			c += 1
		}
	}

	if c == 2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	packages.Greet()
	util.Hello()

}

//output:
/*
value of a is : 10
 b value is 50.7

 name is ,new
 bool :false0020
2510 20 30
 string          10
3.14
 23 hyd
*/
