package main

import "fmt"

func main() {
	var age int
	fmt.Println("Enter age")
	fmt.Scan(&age)
	result, err := divide(10, 0)
	res := checkAge(age)
	number := test(age)
	if number != nil {
		fmt.Println(number)
	} else {
		fmt.Println("positive")
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	if res != nil {
		fmt.Println(res)
	} else {
		fmt.Println("eligible for vote")
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	panic("error occurred")
}

//basic error handling
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// using fmt.Errorf : fmt.Errorf allows formatted error messages
func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age %d is below 18 ,not eligible for vote", age)
	}
	return nil
}

//custom Error handling
type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return e.msg
}

func test(n int) error {
	if n < 0 {
		return MyError{"Negative number"}
	}
	return nil
}
