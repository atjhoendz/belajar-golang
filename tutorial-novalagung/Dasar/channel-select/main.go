package main

import (
	"fmt"
	"runtime"
)

func getAverage(numbers []int, ch chan float64) {
	var sum int = 0

	for _, num := range numbers {
		sum += num
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max int = numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	ch <- max
}

func main() {
	runtime.GOMAXPROCS(2)

	var numbers = []int{3, 4, 3, 5, 2, 7, 8, 5, 2, 3, 2, 4}
	fmt.Println("numbers\t: ", numbers)

	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)

	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %d \n", max)
		}
	}
}
