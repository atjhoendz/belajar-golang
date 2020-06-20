package main

import "fmt"

func main() {
	/*
	* deklarasi variable manifest typing (deklarasi dengan menuliskan tipe datanya)
	*
	* var namaVariable tipe-data = value
	* var namaVariable tipe-data
	 */
	var firstName string = "Mohamad"

	var lastName string
	lastName = "Armando"

	fmt.Printf("Halo %s %s\n", firstName, lastName) // Halo Mohamad Armando

	/*
	* deklarasi variable type inference
	* deklarasi tanpa menuliskan tipe data
	* otomatis mengetahui tipe datanya berdasarkan value pertama yang diassign
	*
	* namaVariable := value
	* := hanya digunakan 1 kali pada saat deklarasi awal
	* selanjutnya cukup = saja untuk re-assign
	 */
	middleName := "Achu"
	middleName = "Achun"

	fmt.Println(firstName, middleName, lastName) // Mohamad Achun Armando

	/*
	* deklarasi multi variable
	* mendeklarasikan banyak variable sekaligus dalam 1 line
	 */
	var nama1, nama2, nama3 string = "Mohamad", "Achun", "Armando"

	fmt.Println(nama1, nama2, nama3) // Mohamad Achun Armando

	var angka1, angka2, angka3 int
	angka1, angka2, angka3 = 1, 2, 3

	fmt.Println(angka1, angka2, angka3) // 1 2 3

	buah1, angka4, nama4 := "Jeruk", 4, "Caesarini"

	fmt.Println(buah1, angka4, nama4) // Jeruk 4 Caesarini

	/*
	* variable _ (underscore) (reserved variable)
	* merupakan variable tong sampah
	* untuk menampung nilai yang tidak dipakai
	* untuk menampung nilai balik dari fungsi yang tidak digunakan
	* tidak dapat ditampilkan valuenya
	* cukup dengan = untuk assign value
	 */
	_ = "sampah"
	_ = "makanan"
	nama, _ := "Jack", "ga dipake"

	fmt.Println(nama) // Jack

	/*
	* deklarasi variable menggunakan new
	* digunakan untuk membuat variable pointer
	* berisikan alamat dari variable
	* untuk mengakses valuenya menggunakan tanda asterik (*)
	 */
	hewan := new(string)

	fmt.Println(hewan)  // 0xc0000102b0
	fmt.Println(*hewan) // ""

	*hewan = "ini hewan"
	fmt.Println(*hewan) // ini hewan

}
