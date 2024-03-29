package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go print(5, "this is goroutine")
	print(5, "this is main goroutine")

	var input string
	fmt.Scanln(&input)
}
