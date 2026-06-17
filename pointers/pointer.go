package main

import "fmt"

func checkVal(a *int) { // *int means it holds address of a
	*a = 50 //here * holds the value of a
}

type Person struct {
	name string
	age  int
}

func Update(p *Person) {
	p.name = "ravi"
	p.age = 20
}

//sharing the data btw functions
func add(x *int) {
	*x += 10
}

func mul(x *int) {
	*x *= 2
}

func main() {
	var x = 10
	b := 10
	c := 2
	checkVal(&b)
	fmt.Println("b", b)
	fmt.Println("x", x)
	fmt.Println("address of x is", &x)
	var p1 = Person{"ravi", 25}
	Update(&p1)
	fmt.Println(p1)

	add(&c)
	mul(&c)
	fmt.Println(c)

	// Method 1: declare then assign
	var u1 User
	u1.ID = 1
	u1.Name = "Alice"
	u1.Email = "alice@mail.com"
	u1.IsActive = true

	// Method 2: struct literal
	u2 := User{
		ID:       2,
		Name:     "Bob",
		Email:    "bob@mail.com",
		IsActive: false,
	}

	// Method 3: short, order matters
	u3 := User{3, "Charlie", "charlie@mail.com", true}

	fmt.Println(u1, u2, u3)
}

// define a struct
type User struct {
	ID       int
	Name     string
	Email    string
	IsActive bool
}
