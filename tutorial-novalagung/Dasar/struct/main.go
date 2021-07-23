package main

import "fmt"

// deklarasi struct
type student struct {
	name  string
	grade int
}

/*
 * embedded struct
 * memasukan struct menjadi properti pada struct lain
 *
 * dapat memiliki properti yang sama dengan struct parentnya
 * untuk pemanggila properti yang harus spesifik
 * misal jika properti yang sama pada studentB adalah age maka: s2.person.age
 */
type person struct {
	name string
	age  int
}

type studentB struct {
	grade int
	person
}

func main() {
	/*
	 * struct merupakan tipe data yang berisikan kumpulan variable dengan berbagai tipe data
	 * dan atau method(fungsi) didalamnya
	 */

	// inisialisasi struct student
	var s1 student
	s1.name = "achun"
	s1.grade = 100

	s2 := student{}
	s2.name = "delanika"
	s2.grade = 80

	s3 := student{"mohamad", 80}

	s4 := student{name: "armando", grade: 90}

	fmt.Println("S1")
	printMessage(s1.name, s1.grade)

	fmt.Println("\nS2")
	printMessage(s2.name, s2.grade)

	fmt.Println("\nS3")
	printMessage(s3.name, s3.grade)

	fmt.Println("\nS4")
	printMessage(s4.name, s4.grade)

	// tipe pointer
	var s5 *student = &s1
	s5.name = "Achun"

	fmt.Println("\nS5")
	printMessage(s5.name, s5.grade)

	// mengakses embedded struct
	s6 := studentB{}
	s6.name = "Moh Achun Armando"
	s6.age = 20
	s6.grade = 100

	fmt.Println("\nS6")
	fmt.Println("Name:", s6.name)
	fmt.Println("Age:", s6.age)
	fmt.Println("Age:", s6.person.age)
	fmt.Println("Grade:", s6.grade)

	/*
	 * anonymous struct
	 * struct yang tidak dideklarasi diawal sebagai tipe data baru,
	 * lansung ketika pembuatan objek
	 */
	s7 := struct {
		person
		grade int
	}{}

	s7.name = "Armando"
	s7.age = 21
	s7.grade = 98

	fmt.Println("\nS7")
	fmt.Println("Name:", s7.name)
	fmt.Println("Age:", s7.age)
	fmt.Println("Age:", s7.person.age)
	fmt.Println("Grade:", s7.grade)

	/*
	 * kombinasi anonymous struct dengan slice
	 *
	 */
	allPerson := []person{
		{name: "Mohamad", age: 20},
		{name: "Achun", age: 22},
		{name: "Armando", age: 21},
	}

	fmt.Println()
	for _, v := range allPerson {
		fmt.Println("Nama:", v.name)
		fmt.Println("Age:", v.age)
		fmt.Println()
	}

	/*
	 * pembuatan anonymous struct dengan var
	 */
	var fruits = struct {
		name string
		qty  int
	}{
		name: "Apel", qty: 10,
	}

	fmt.Println(fruits.name, "with qty:", fruits.qty)

	/*
	 * nested struct
	 */
	type studentC struct {
		person struct {
			name string
			age  int
		}
		grade   int
		hobbies []string
	}
}

func printMessage(name string, grade int) {
	fmt.Println("Name:", name)
	fmt.Println("Grade:", grade)
}
