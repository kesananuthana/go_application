package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3) //it means we are starting 2 goroutines
	go task("A", &wg)
	go task("B", &wg)
	task("c", &wg)
	go hello()
	wg.Wait() //wait until finish
	fmt.Println("All tasks are completed")
	go display()
	say()
	time.Sleep(time.Second)
}

func say() {
	fmt.Println("Hello from Goroutine")
}

func hello() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}

func display() {
	for i := 0; i < 3; i++ {
		fmt.Println("display hii")
	}
}

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done() // tell "I am finished"
	fmt.Println("I am ", name)
}

/* this is time.sleep code
func main() {
	go hello()
	go display()
	time.Sleep(time.Second)
}

func say() {
	fmt.Println("Hello from Goroutine")
}

func hello() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}

func display() {
	for i := 0; i < 3; i++ {
		fmt.Println("display hii")
	}
}
*/

//wait group function
/*func main() {
	var wg sync.WaitGroup
	wg.Add(3) //it means we are starting 2 goroutines
	task("A", &wg)
	task("B", &wg)
	task("c", &wg)

	wg.Wait() //wait until finish
	fmt.Println("All tasks are completed")
}

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done() // tell "I am finished"
	fmt.Println("I am ", name)
}*/
