package main

import (
	"fmt"
	"strings"
)

/*
 * method
 * merupakan fungsi yang berada pada struct atau tipe data lainnya
 *
 * method berguna pada proses enkapsulasi, untuk mengakses properti private
 *
 * func (insialisasi struct) namaMethod(){}
 */

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("Hello", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

// method pointer
func (s student) changeName1(name string) {
	fmt.Println("--> on changeName1, name changed to:", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("--> on changeNam2, name changed to:", name)
	s.name = name
}

func main() {
	s1 := student{name: "Mohamad Achun Armando", grade: 100}

	s1.sayHello()
	fmt.Println("Get Name at 1:", s1.getNameAt(2))

	fmt.Println()

	s1.changeName1("Armando")
	fmt.Println(s1.name) // name hanya berubah didalam objek changeName1 saja

	s1.changeName2("Atjhoendz")
	fmt.Println(s1.name)
}
