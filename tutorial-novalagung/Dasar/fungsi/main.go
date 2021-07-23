package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {

	// fungsi merupakan sekumpulan blok kode dengan nama tertentu

	// penggunaan fungsi printMessage
	arrName := []string{"Mohamad", "Achun", "Armando"}
	printMessage("Halo my name is", arrName)

	// penggunaan fungsi randomNumber

	// -> untuk membuat fungsi rand benar-benar acak berdasarkan unix time
	rand.Seed(time.Now().Unix())

	randomValue := randomWithRange(3, 10)
	fmt.Println("random 1:", randomValue)
	randomValue = randomWithRange(10, 15)
	fmt.Println("random 2:", randomValue)
	randomValue = randomWithRange(1, 5)
	fmt.Println("random 3:", randomValue)

	// penggunaan fungsi divideNumber
	divideNumber(10, 2)
	divideNumber(1, 0)
	divideNumber(16, -4)

	// penggunaan fungsi calculate
	var diameter float64 = 15
	area, circumference := calculate(diameter)

	fmt.Printf("Luas Lingkaran \t\t: %.2f\n", area)
	fmt.Printf("Keliling Lingkaran \t: %.2f\n", circumference)

	// penggunaan fungsi calculate2
	area2, cicircumference2 := calculate2(diameter)

	fmt.Printf("Luas lingkaran \t\t: %.2f\n", area2)
	fmt.Printf("Keliling lingkaran \t: %.2f\n", cicircumference2)

	// penggunaan fungsi rata-rata
	rataRata := hitungRataRata(9, 10, 5, 6, 9, 3, 6, 7, 8, 5)

	// -> sprintf() digunakan untuk konversi menjadi string, harus dimasukkan kedalam variable
	msg := fmt.Sprintf("Hasil rata-rata: %.2f", rataRata)

	fmt.Println(msg)

	// mengisi parameter fungsi variadic menggunakan slice
	nilai := []int{10, 3, 6, 9, 10, 6, 8, 5, 6, 7}

	// -> menggunakan ... disebelah kanan variable (variable...)
	rataRata = hitungRataRata(nilai...)

	// penggunaan fungsi yourHobbies
	myHobbies := []string{"ngoding", "workout", "eat"}
	yourHobbies("Achun", myHobbies...)

	/*
	 * fungsi closure
	 * sebuah fungsi yang dimasukan kedalam variable
	 *
	 * dapat digunakan untuk membuat fungsi didalam fungsi
	 * membuat fungsi mengembalikan fungsi
	 *
	 * ditempatkan didalam fungsi
	 */

	getMinMax := func(n []int) (min int, max int) {
		for i, e := range n {
			switch {
			case i == 0:
				min, max = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}

		return
	}

	numbers := []int{3, 4, 5, 7, 1, 5, 8, 4, 9, 0}
	min, max := getMinMax(numbers)

	fmt.Printf("Daftar Angka: %v\n", numbers)
	fmt.Printf("Angka terkecil: %d\n", min)
	fmt.Printf("Angka terbesar: %d\n", max)

	/*
	 * immediately-invoked function expression (IIFE)
	 * fungsi yang langsund dipanggil setelah dideklarasikan
	 */
	dataNumbers := []int{1, 4, 6, 3, 7, 4, 7, 9, 0, 5, 2, 6, 7}

	filteredNumbers := func(min int) []int {
		var result []int

		for _, v := range dataNumbers {
			if v < min {
				continue
			}
			result = append(result, v)
		}

		return result
	}(4) // memasukan argument (4) diakhir fungsi

	fmt.Println("\nNumbers:", numbers)
	fmt.Println("Filtered Numbers:", filteredNumbers)

	// penggunaan fungsi closure sebagai return value pada fungsi findMax()
	newMax := 5
	anotherNumbers := []int{4, 2, 5, 6, 7, 3, 5, 0, 8, 4, 6}
	totalNumber, newFilteredNumbers := findMax(anotherNumbers, newMax)

	fmt.Println("\nNumbers:", anotherNumbers)
	fmt.Println("Max:", newMax)

	fmt.Println("Total Result:", totalNumber)
	fmt.Println("Filtered Numbers:", newFilteredNumbers())

	// penggunaan fungsi closure sebagai parameter pada fungsi filter()
	data := []string{"Mohamad", "Achun", "Armando"}

	dataContainsH := filter(data, func(each string) bool {
		return strings.Contains(each, "h")
	})

	dataLength5 := filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println()
	fmt.Println("Data:", data)
	fmt.Println("Data dengan huruf h:", dataContainsH)
	fmt.Println("Data dengan panjang 5:", dataLength5)
}

/*
 * fungsi void
 *
 * func namaFungsi(param1 tipeData, param2 tipeData) {}
 * atau
 * jika tipe data parameter sama
 * func namaFungsi(param1, param2, tipeData) {}
 */

func printMessage(message string, arr []string) {
	nameString := strings.Join(arr, " ")

	fmt.Println(message, nameString)
}

/*
 * fungsi return value
 *
 * func namaFungsi(param1 tipeData1, param2 tipeData2) tipeDataOutput {}
 * atau
 * jika tipe data parameter sama
 * func namaFungsi(param1, param2, tipeDataYangSama) tipeDataOutput {}
 */
func randomWithRange(min, max int) int {
	return rand.Int()%(max-min+1) + min
}

/*
 * penggunaan return
 * untuk menghentikan proses didalam fungsi
 *
 */
func divideNumber(number1, number2 int) {
	if number2 == 0 {
		fmt.Printf("Tidak dapat membagi dengan %d\n", number2)
		return
	}

	result := number1 / number2
	fmt.Printf("Hasil %d/%d = %d\n", number1, number2, result)
}

/*
 * fungsi multiple return value
 *
 * func namaFungsi(parameter) (output1, output2)
 */

func calculate(d float64) (float64, float64) {
	area := math.Pi * math.Pow(d/2, 2)

	circumference := math.Pi * d

	return area, circumference
}

/*
 * predefined return value
 *
 * func namaFungsi(varParam tipedataParam)(varReturn1 tipedataReturn1, varReturn2 tipedataReturn2)
 */
func calculate2(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)

	circumference = math.Pi * d

	return
}

/*
 * fungsi variadic
 * merupakan fungsi dengan argumen yang tidak terbatas yang tipe datanya sama
 * yang dimasukkan hanya kepada 1 variable dalam 1 parameter
 *
 * func namaFungsi(variable ...tipedata)
 */
func hitungRataRata(numbers ...int) float64 {
	total := 0
	for _, number := range numbers {
		total += number
	}

	// float64(variable) digunakan untuk mengcasting tipe data menjadi float64
	avg := float64(total) / float64(len(numbers))

	return avg
}

// penggabungan fungsi dengan parameter biasa & variadic
func yourHobbies(name string, hobbies ...string) {
	hobbiesAsString := strings.Join(hobbies, " ")

	fmt.Printf("Your Name: %s\n", name)
	fmt.Printf("Your Hobbies: %s\n", hobbiesAsString)
}

/*
 * fungsi closure sebagai return value
 */
func findMax(numbers []int, max int) (int, func() []int) {
	res := []int{}

	for _, v := range numbers {
		if v <= max {
			res = append(res, v)
		}
	}

	return len(res), func() []int {
		return res
	}
}

/*
 * fungsi closure sebagai parameter
 * atau bisa disebut fungsi callback
 *
 * fungsi filter ini akan mengecek data berdasarkan fungsi yang ada didalam fungsi callback
 */
func filter(data []string, callback func(string) bool) []string {
	var result []string

	for _, each := range data {
		// memanggil fungsi callback untuk mengecek setiap data yang ada
		if filtered := callback(each); filtered { // assign value filtered dari hasil callback, kemudia cek apakah true/false
			result = append(result, each)
		}
	}
	return result
}

/*
 * alias pada fungsi closure
 *
 * type
 *
 * type namaAlias func(tipedata) tipeData
 *
 * penggunaan
 * type FilterCallback func(string) bool
 *
 * func filter(data []string, callback FilterCallback) []string {}
 */
