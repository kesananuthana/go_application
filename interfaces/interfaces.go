package main

import "fmt"

//Single method interfaces

/*type Game interface {
	NoOfPalyers() int

}

type Volleyball struct {
	p int
}

func (v Volleyball) NoOfPalyers() int {
	return v.p
}

func main() {
	var g Game
	a := Volleyball{6} //This is just creating a struct value (an object).
	g = a
	fmt.Println(g.NoOfPalyers())
}*/

//Multiple method interfaces
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	l, b float64
}

type Square struct {
	a float64
}

func (r Rectangle) Area() float64 {
	return r.l * r.b
}

func (r1 Rectangle) Perimeter() float64 {
	return 2 * (r1.l + r1.b)
}

func (s Square) Area() float64 {
	return s.a * s.a
}

func (s1 Square) Perimeter() float64 {
	return 4 * s1.a
}

//type switch
func checkType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Println("integer", v)
	case string:
		fmt.Println("string", v)
	case float64:
		fmt.Println("float", v)
	default:
		fmt.Println("unknown type")
	}
}

func printValue(value any) {
	fmt.Println(value)
}

func main() {
	var l, b, a float64
	fmt.Println("Enter the values of l and b")
	fmt.Scan(&l, &b)
	fmt.Println("Enter square values")
	fmt.Scan(&a)

	r := Rectangle{l, b}
	s := Square{a}
	shapes := []Shape{r, s} //Create a slice of type Shape and put r and s inside it
	for _, i := range shapes {
		fmt.Println(i.Area())
		fmt.Println(i.Perimeter())
	}

	//empty interfaces
	//Empty interface is used for : flexibility when type is unknown
	var x interface{} //Since there are no methods to implement, any value can be stored in x.In Go, an empty interface is an interface that has no methods.
	x = 10
	fmt.Println(x)
	x = "hello"
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)

	checkType("hello")

	//slice of empty interface
	example := []interface{}{2, 3.5, "hello"}
	for _, val := range example {
		printValue(val)
	}
}
