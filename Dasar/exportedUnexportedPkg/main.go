package main

import (
	"exportedUnexportedPkg/library"
	"exportedUnexportedPkg/student"
	f "fmt"
)

/*
 * Exported atau public: komponen dapat diakses oleh package lain
 * Unexported atau private: komponen hanya dapat diakses dalam package yang sama atau folder yang sama
 *
 * ----------- IMPORTANT ---------------
 * Penulisan modifier => Character Case
 * pada penamaan fungsi, struct, variable
 * -------------------------------------
 *
 * Exported: huruf awal besar
 * Unexported: huruf awal kecil
 *
 * penggunaan . (titik)
 * import . "exportedblalbl/library"
 * -> untuk membuat package yang diimport satu level dengan peng-import
 * -> jadi untuk mengakses method/struct/properti dapat langsung tanpa nama packagenya
 * -> cth: Student{"Atjhoendz", 100}
 *
 * alias
 * import f "fmt"
 * f.Println()
 */
func main() {
	s1 := library.Student{"atjhoendz", 100}

	f.Println("Name:", s1.Name)
	f.Println("Grade:", s1.Grade)

	/*
	 * fungsi yang ada pada partial.go
	 *
	 * untuk mengeksekusinya
	 * go run main.go partial.go
	 */
	// sayPartial("partial")

	f.Println("\nstudent.go")
	f.Println("Name:", student.Student.Name)
	f.Println("Grade:", student.Student.Grade)
}
