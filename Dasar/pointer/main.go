package main

import "fmt"

func main() {
	/*
	 * pointer
	 * -> merupakan alamat memori dimana suatu data disimpan
	 *
	 * (&variable) untuk mengambil alamat
	 * (*variable) untuk mengambil value
	 *
	 * deklarasi
	 * var variable *tipedata
	 */
	nomorA := 4
	var nomorB *int = &nomorA

	fmt.Println("Value dari nomorA\t:", nomorA)
	fmt.Println("Alamat dari nomorA\t:", &nomorA)

	fmt.Println("\nValue dari nomorB\t:", nomorB)
	fmt.Println("Alamat dari nomorB\t:", &nomorB)
	fmt.Println("Value dari alamat yang ditunjuk oleh nomorB\t:", *nomorB)

	// perubahan pada variable yang direference akan mengubah nilai pada variable pointer juga
	nomorA = 5

	fmt.Println()
	fmt.Println("Value dari nomorA\t:", nomorA)
	fmt.Println("Alamat dari nomorA\t:", &nomorA)

	fmt.Println("\nValue dari nomorB\t:", nomorB)
	fmt.Println("Alamat dari nomorB\t:", &nomorB)
	fmt.Println("Value dari alamat yang ditunjuk oleh nomorB\t:", *nomorB)

	// penggunaan pointer pada parameter pada fungsi change()
	number := 4
	fmt.Println("\nNumber before:", number)

	change(&number, 10)
	fmt.Println("Number after:", number)
}

func change(original *int, value int) {
	*original = value
}
