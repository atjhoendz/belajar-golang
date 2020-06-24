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

	// untuk membuat fungsi rand benar-benar acak berdasarkan unix time
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
