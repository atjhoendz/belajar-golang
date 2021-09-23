package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Armando")

	fmt.Println(message)
}
