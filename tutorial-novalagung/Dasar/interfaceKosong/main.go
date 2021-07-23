package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

/*
 * Interface kosong -> interface{}
 * merupakan tipe data yang data diisi segala jenis tipe data
 * biasa disebut dynamic typing
 */

func main() {
	var secret interface{}

	secret = "atjhoend armando"
	fmt.Println(secret)

	secret = []string{"apel", "mangga", "pisang"}
	fmt.Println(secret)

	secret = 22.3
	fmt.Println(secret)

	data := map[string]interface{}{
		"buah":    "apel",
		"nomor":   1,
		"pi":      3.14,
		"deretan": []int{1, 3, 4, 5, 3, 4},
	}
	fmt.Println(data)

	// casting tipe data interface{}
	var data2 interface{}
	data2 = 2

	// untuk melakukan operasi maka harus dicasting ke tipe data aslinya, jika tidak maka error
	// type assertions
	result := data2.(int) * 10
	fmt.Println("2x10 = ", result)

	data2 = []string{"apel", "mangga", "jeruk"}
	fruits := strings.Join(data2.([]string), ", ")

	fmt.Println(fruits, "is my favorites fruits")

	// casting interface{} ke objek pointer
	var data3 interface{} = &person{"dela", 20}
	name := data3.(*person).name

	fmt.Println(name)

	// kombinasi slice, map dan interface{}
	personMap := []map[string]interface{}{
		{"name": "Mohamad", "age": 20},
		{"name": "Achun", "age": 20},
		{"name": "Armando", "age": 24},
	}

	for _, each := range personMap {
		fmt.Println(each["name"], "age is", each["age"])
	}

	fruits2 := []interface{}{
		map[string]interface{}{"name": "atjhoendz", "age": 20},
		[]string{"jeruk", "mangga", "apel"},
		"pepaya",
	}

	fmt.Println()

	for _, each := range fruits2 {
		fmt.Println(each)
	}
}
