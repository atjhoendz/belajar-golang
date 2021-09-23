package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Mohamad", "Achun", "Armando"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	fmt.Println("-------")

	for i, message := range message {
		fmt.Println(i)       // print key
		fmt.Println(message) // print value
	}
}
