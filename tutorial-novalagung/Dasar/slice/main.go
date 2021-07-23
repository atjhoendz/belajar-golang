package main

import "fmt"

func main() {
	/*
	* slice merupakan data reference dari array
	* dapat dibuat dari manipulasi array atau slice lain
	* perubahan data dari suatu slice akan berdampak ke slice lain (reference)
	*
	* var slice = []string{data}
	*
	* perbedaan mencolok slice dengen array terletak pada deklarasi jumlah datanya
	* slice tidak didefinisikan
	* array harus didefinisikan
	 */
	var sliceBuah = []string{"apel", "jeruk", "pisang", "kelapa"}

	fmt.Println("data slice:", sliceBuah)
	fmt.Println("data ke 1:", sliceBuah[0])

	//	pengaksesan data dengan 2 indeks

	fmt.Println(sliceBuah[0:4]) // indeks 0 hingga sebelum 4 => [apel, jeruk, pisang, kelapa]
	fmt.Println(sliceBuah[1:4]) // index 1 hingga sebelum 4 => [jeruk, pisang, kelapa]
	fmt.Println(sliceBuah[0:0]) // index 0 hingga sebelum 0 => [] (tidak ada nilai dengan indeks sebelum 0)
	fmt.Println(sliceBuah[:])   // ambil semua data => [apel, jeruk, pisang, kelapa]
	fmt.Println(sliceBuah[2:])  // ambil semua data dimulai dari indeks ke 2 => [pisang, kelapa]
	fmt.Println(sliceBuah[:2])  // ambil semua data hingga sebelum indeks ke 2 => [apel, jeruk]

	/*
	* karena slice merupakan reference
	*
	* ketika membuat slice baru dari sebuah slice
	* perubahan data yang terjadi pada slice baru akan turut mengubah slice lama
	 */
	namaLengkap := []string{"Mohamad", "Achun", "Armando"}

	// pembuatan slice baru menggunakan 2 indeks
	namaDepan := namaLengkap[0:1]
	namaTengah := namaLengkap[1:2]
	namaBelakang := namaLengkap[2:3]

	fmt.Println(namaLengkap)
	fmt.Println(namaDepan)
	fmt.Println(namaTengah)
	fmt.Println(namaBelakang)

	namaTengah[0] = "Delanika"

	fmt.Println(namaLengkap)
	fmt.Println(namaDepan)
	fmt.Println(namaTengah)
	fmt.Println(namaBelakang)

	// fungsi untuk mengoperasikan slice
	// len(slice) => untuk menghitung jumlah elemen dari slice
	fmt.Println("Jumlah elemen:", len(namaLengkap))

	// cap(slice) => untuk menghitung lebar dan kapasitas dari slice
	fmt.Println(cap(namaLengkap))

	// append() => untuk menambah elemen pada slice
	fruits := []string{"apel", "mangga"}

	newFruits := append(fruits, "semangka")

	fmt.Println(fruits)
	fmt.Println(newFruits)

	/*
	* catatan append
	* jika len(newSlice) == cap(newSlice) maka menambah referensi baru
	* jika len(newSlice) < cap(newSlice) maka elemen baru ditempatkan pada cakupan cap,
	*		sehingga slice lain dengan referensi yang sama akan ikut berubah
	 */

	fruits2 := []string{"apel", "mangga", "pisang"}

	bfruits := fruits2[0:2]

	fmt.Println(cap(bfruits)) // 3
	fmt.Println(len(bfruits)) // 2

	fmt.Println(fruits2) // [apel, mangga, pisang]
	fmt.Println(bfruits) // [apel, mangga]

	cfruits := append(bfruits, "nanas")

	fmt.Println(fruits2) // [apel, mangga, nanas]
	fmt.Println(bfruits) // [apel, mangga]
	fmt.Println(cfruits) // [apel, mangga, nanas]

	/*
	* copy() => untuk mengcopy elemen slice dari src ke dst
	*
	* copy(dst, src)
	 */

	dst := make([]string, 3)

	src := []string{"apel", "mangga", "pisang"}

	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	// mengakses slice dengan 3 indeks
	// fruits[0:1:2] => 2 dibelakang adalah kapasitas dari slice yang baru tidak boleh lebih dari kapasitas slice lama

	fruits3 := []string{"apel", "mangga", "nanas"}
	dfruits := fruits3[0:2]
	efruits := fruits3[0:2:2]

	fmt.Println(fruits3)
	fmt.Println(cap(fruits3)) // 3
	fmt.Println(len(fruits3)) // 3

	fmt.Println(dfruits)
	fmt.Println(cap(dfruits)) // 3
	fmt.Println(len(dfruits)) // 2

	fmt.Println(efruits)
	fmt.Println(cap(efruits)) // 2
	fmt.Println(len(efruits)) // 2
}
