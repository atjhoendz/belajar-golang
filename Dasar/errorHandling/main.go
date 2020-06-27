package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
 * error merupakan suatu tipe data
 */

func main() {
	var input string
	fmt.Print("masukan nomor: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}

	// penggunaan fungsi validate()
	var name string
	fmt.Print("Masukan nama: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Halo", name)
	} else {
		fmt.Println(err.Error())
	}

	/* penggunaan panic()
	 * panic digunakan untuk menampilkan stack trace error, sekaligus menghentikan flow goroutine
	 * main() termasuk goroutine
	 */

	var nameB string
	fmt.Print("Masukan nama B: ")
	fmt.Scanln(&nameB)

	if valid, err := validate(nameB); valid {
		fmt.Println("Halo", nameB)
	} else {
		panic(err.Error())
		// fmt.Println("End") // code apapun setelah panic tidak akan dieksekusi
	}

	/* penggunaan recover()
	 * digunakan untuk menghandle panic error agar tidak menampilkan banyak message
	 */
	defer catch() // defer digunakan untuk memanggil fungsi pada saat suatu blok fungsi akan berakhir

	var nameC string
	fmt.Print("Masukan nama C: ")
	fmt.Scanln(&nameC)

	if valid, err := validate(nameC); valid {
		fmt.Println("Halo", nameC)
	} else {
		panic(err.Error())
		// fmt.Println("End") // code apapun setelah panic tidak akan dieksekusi
	}

	// penerapan recover pada IIFE
	data := []string{"superman", "batman", "wonder woman"}

	for _, each := range data {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping:", each, "| message:", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()
			panic("some error hapen")
		}()
	}
}

// pembuatan custom error
func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

// penggunaan recover
func catch() {
	if r := recover(); r != nil {
		fmt.Println("error occured: ", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}
