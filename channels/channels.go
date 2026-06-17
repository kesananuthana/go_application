package main

import "fmt"

/*func main() {
	ch := make(chan string) // it means creating the channel

	go func() {
		ch <- "hello channel" // sends data into channel
	}()
	msg := <-ch //take data from channel
	fmt.Println(msg)
}
*/

//channels can also works on normal functions
//we can also create multiple channels
func main() {
	ch := make(chan string, 1)
	ch1 := make(chan int, 2)
	ch1 <- 5
	ch1 <- 10
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	task(ch)
	msg := <-ch
	fmt.Println(msg)
}

func task(ch chan string) {
	ch <- "Normal function"
}
