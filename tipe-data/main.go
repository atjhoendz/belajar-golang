package main

import "fmt"

func main() {
	/*
	* tipe data numerical
	* uint (unsigned / positif)
	* int (signed / bilangan bulat)
	* byte = int8
	* rune = int32
	 */
	var positiveNumber uint = 19
	var negativeNumber = -10 // tipe data otomatis

	fmt.Println(positiveNumber, negativeNumber)         // 19 -10
	fmt.Printf("Positive Number: %d\n", positiveNumber) // Positive Number: 19
	fmt.Printf("Negative Number: %d\n", negativeNumber) // Negative Number: -10

	/*
	* tipe data decimal
	* floating point
	 */
	var decimal float32 = 8.10
	decimal2 := 3.14

	fmt.Printf("Angka desimal: %f\n", decimal)    // 8.100000
	fmt.Printf("Angka desimal: %.2f\n", decimal2) // 3.14

	/*
	* tipe data boolean
	* true dan false
	 */
	var benar bool = true
	fmt.Printf("Apakah benar? %t \n", benar) // Apakah benar? true

	/*
	* tipe data string
	 */
	var nama string = "Mohamad Achun Armando"
	string2 := `
		ini adalah backticks
		value sesuai deklarasi
	`
	fmt.Printf("Nama: %s\n", nama) // Nama: Mohamad Achun Armando
	fmt.Printf("string2: %s\n", string2)

	/*
	* Nilai Nil & Zero Value
	* nil => null => nilai kosong
	* zero value => nilai default suatu tipe data
	*
	* zero value string => ""
	* zero value bool => false
	* zero value tipe numerik non desimal => 0
	* zero value tipe numerik desimal => 0.0
	*
	* nil hanya bisa diassign pada tipe variable selain diatas, yaitu:
	* - pointer
	* - tipe data fungsi
	* - slice, map, channel
	* - interface kosong => interface{}
	 */
	var zeroValueInt int
	nilaiNil := new(string)
	nilaiNil = nil

	fmt.Println(zeroValueInt) // 0
	fmt.Println(nilaiNil)     // <nil>
}
