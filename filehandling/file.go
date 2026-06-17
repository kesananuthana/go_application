package main

import (
	"fmt"
	"os"
)

//file creation
/*func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	fmt.Println("File created successfully")
}*/

// data insertion
/*func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	data := "file handling functions"
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("data inserted successfully")
	defer file.Close()

}*/

// reading data from a file
func main() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("data :", string(data))
	}
}
