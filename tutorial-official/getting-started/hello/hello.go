package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Printf("\nQuote: %s\n", quote.Go())
}
