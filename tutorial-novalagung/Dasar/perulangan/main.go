package main

import "fmt"

func main() {
	// pada go perulangan hanya menggunakan for

	// perulangan for biasa
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for dengan argumen hanya kondisi
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}

	// for tanpa argumen
	a := 0

	for {
		fmt.Println(a)

		a++

		if a == 5 {
			break
		}
	}

	// penggunaan break dan continue
	for i := 0; i <= 10; i++ {
		if i%2 == 1 {
			continue //untuk melanjutkan perulangan ke iterasi selanjutnya
		}

		if i > 8 {
			break //untuk menghentikan perulangan
		}

		fmt.Println("Angka", i)
	}

	// perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	// pemanfaatan label pada perulangan

outerLoop:
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if i == 3 {
				break outerLoop // untuk memberhentikan semua perulangan
			}
			fmt.Printf("matriks [%d][%d]\n", i, j)
		}
	}

	/*
	* for range
	*
	* for i, name := range names
	* for _, name := range names (jika indeks tidak dibutuhkan)
	* for i, _ := range names (jika value tidak dibutuhkan)
	* for i := range names (jika value tidak dibutuhkan)
	 */
	var names = [...]string{"mohamad", "achun", "armando"}

	for i, name := range names {
		fmt.Printf("Element ke %d: %s\n", i, name)
	}
}
