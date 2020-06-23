package main

import "fmt"

func main() {
	/*
	* map merupakan tipe data asosiatif
	* berbentuk key-value pair
	 */
	fruits := map[string]int{}

	fruits["apel"] = 40
	fruits["nanas"] = 54
	fruits["pisang"] = 34

	fmt.Println("Apel:", fruits["apel"])
	fmt.Println("Nanas:", fruits["nanas"])
	fmt.Println("Pisang:", fruits["pisang"])

	// cara lain inisialisasi

	books := map[string]int{
		"naruto":    1,
		"doraemaon": 4,
	}

	fmt.Println(books)

	// ------
	map1 := make(map[string]int)
	map2 := *new(map[string]int)

	fmt.Println(map1, map2)

	// iterasi map menggunakan for-range
	drinks := map[string]int{
		"coklat":    23,
		"teh manis": 3,
		"es jeruk":  4,
		"josu":      5,
	}

	for key, value := range drinks {
		fmt.Println(key, ":", value)
	}

	// menghapus data pada map
	// delete(namaMap, "key")
	newDrinks := map[string]int{
		"josu":   2,
		"apel":   1,
		"nurdin": 4,
	}

	fmt.Println(newDrinks)

	delete(newDrinks, "apel")

	fmt.Println(newDrinks)

	// mencari value berdasarkan key
	fmt.Println(drinks)
	value, isExist := drinks["coklat"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	/*
	* kombinasi map dengan slice
	*
	* []map[string]string{}
	*
	 */
	dataPesanan := []map[string]int{
		{"josu": 5, "nurdin": 10},
		{"omlet": 3, "telorbaso": 4},
		{"kerupuk": 1},
	}

	for _, data := range dataPesanan {
		fmt.Println(data)
	}
}
