package main

import "fmt"

func main() {
	// array merupakan kumpulan data yang bertipe sama yang memiliki indeks dimulai dari 0
	var buah [4]string

	buah[0] = "apel"
	buah[1] = "belimbing"
	buah[2] = "ceri"
	buah[3] = "durian"

	fmt.Println(buah[0], buah[1], buah[2], buah[3])

	// deklarasi array langsung assign value
	var binatang = [3]string{"sapi", "monyet", "rusa"}

	fmt.Println("Jumlah array:", len(binatang))
	fmt.Println("Isi element:", binatang)

	// deklarasi array langsung assign value secara vertical
	var orang [3]string

	orang = [3]string{
		"maman",
		"jamal",
		"jack", // koma terakhir harus dituliskan pada deklarasi gaya vertical
	}

	fmt.Println("Data orang:", orang)

	// deklarasi array langsung assign value tanpa deklarasi panjang array
	var nomor = [...]int{1, 2, 3, 5, 4}

	fmt.Println("Jumlah element \t:", len(nomor))
	fmt.Println("Isi element \t:", nomor)

	// array multidimensi
	var matriks = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("matriks:", matriks)

	// perulangan array menggunakan for
	var dataNama = [...]string{"delanika", "olympiani", "trieswari", "caesarini"}

	for i := 0; i < len(dataNama); i++ {
		fmt.Println("Nama", i+1, ":", dataNama[i])
	}

	// perulangan array menggunakan for-range
	var namaLengkap = [...]string{"Mohamad", "Achun", "Armando"}

	for i, nama := range namaLengkap {
		fmt.Printf("Nama ke %d: %s\n", i, nama)
	}

	// perulangan array for-range dengan variable underscore
	var dataBuah = [4]string{"mangga", "apel", "jeruk", "nanas"}

	for _, buah := range dataBuah {
		fmt.Println(buah)
	}

	// alokasi elemen array menggunakan make
	var fruits = make([]string, 3)

	fruits[0] = "jeruk"
	fruits[1] = "semangka"
	fruits[2] = "pepaya"

	fmt.Println("Fruits:", fruits)
}
