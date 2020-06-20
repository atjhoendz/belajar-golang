package main

import "fmt"

func main() {
	/*
	* IF - ELSE
	 */

	nilai := 10

	if nilai >= 8 && nilai <= 10 {
		fmt.Println("Sempurna")
	} else if nilai < 8 && nilai > 6 {
		fmt.Println("Cukup")
	} else {
		fmt.Println("Kurang")
	}

	/*
	* Temporary variable pada IF - ELSE
	* hanya berlaku pada operator kondisi
	* variable bersifat local hanya bisa digunakan didalam blok kondisinya
	* hanya dapat di deklarasi menggunakan :=
	 */
	nilai2 := 9500.0

	if percent := nilai2 / 100; percent >= 80 && percent <= 100 {
		fmt.Printf("%.1f%s sempurna!\n", percent, "%")
	} else if percent > 60 && percent < 80 {
		fmt.Printf("%.1f%s cukup!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s kurang\n", percent, "%")
	}

	/*
	* Switch Case
	*
	* pada statement case atau default dapat menggunakan {}
	* case dapat lebih dari 1 kondisi
	 */
	angka := 1

	switch angka {
	case 5:
		fmt.Println(angka)
	case 4:
		fmt.Println(angka)
	case 1, 2, 3:
		{
			fmt.Println("You are great!")
			fmt.Printf("You are number %d!\n", angka)
		}
	default:
		fmt.Println("Angka tidak ada yang cocok")
	}

	/*
	* Switch Case dengan gaya If-else
	 */
	angka2 := 2

	switch {
	case angka2 > 2:
		fmt.Println(angka2)
	case angka2 > 0 && angka2 <= 2:
		fmt.Println(angka2)
	default:
		fmt.Println("oops...")
	}

	/* Fallthroug pada Switch Case
	* digunakan untuk memaksa proses pengecekan diteruskan ke case selanjutnya
	* meskipun case selanjutnya tersebut bernilai salah
	 */
	angka3 := 3

	switch {
	case angka3 >= 3:
		fmt.Println("mantap")
		fallthrough
	case angka3 < 3:
		fmt.Println("oke")
	default:
		fmt.Println("gapapa")
	}

	/* Seleksi bersarang if-else switch
	 */

	var point2 = 10

	if point2 > 7 {
		switch point2 {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point2 == 5 {
			fmt.Println("not bad")
		} else if point2 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point2 == 0 {
				fmt.Println("try harder!")
			}
		}
	}

}
