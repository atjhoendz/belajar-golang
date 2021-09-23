package main

import "fmt"

func main() {
	var messages = make(chan int, 3)

	go func() {
		for {
			i := <-messages
			fmt.Println("Receive data: ", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Send data: ", i)
		messages <- i
	}
}
