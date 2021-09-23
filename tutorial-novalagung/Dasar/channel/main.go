package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messagesChan = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)

		messagesChan <- data
	}

	go sayHelloTo("first")
	go sayHelloTo("second")
	go sayHelloTo("third")

	var message1 = <-messagesChan
	fmt.Println(message1)

	var message2 = <-messagesChan
	fmt.Println(message2)

	var message3 = <-messagesChan
	fmt.Println(message3)
}
