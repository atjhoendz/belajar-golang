package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

/* mengambil informasi properti dari struct
 * hanya untuk properti public (prefix kapital)
 */
func (s *student) getPropertyInfo() {
	reflectValue := reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	reflectType := reflectValue.Type()

	fmt.Println()
	for i := 0; i < reflectType.NumField(); i++ {
		fmt.Println("nama:", reflectType.Field(i).Name)
		fmt.Println("tipe data:", reflectType.Field(i).Type)
		fmt.Println("nilai:", reflectValue.Field(i).Interface())
		fmt.Println()
	}
}

// mengambil informasi dari method
func (s *student) SetName(name string) {
	s.Name = name
}

/*
 * Reflect merupakan teknik untuk menampilkan informasi dari suatu variable
 *
 * reflect.ValueOf() -> untuk menampilkan value
 * reflect.TypeOf() -> untuk menampilkan tipe datanya
 */

func main() {
	number := 40

	reflectValue := reflect.ValueOf(number)

	fmt.Println("Value reflectValue:", reflectValue)
	fmt.Println("Tipe variable:", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai variable:", reflectValue.Int())
	}

	// mengakses nilai menggunakan interface{}
	number2 := 24
	reflectVal := reflect.ValueOf(number2)

	fmt.Println()
	fmt.Println("Tipe variable:", reflectVal.Type())
	fmt.Println("Value variable:", reflectVal.Interface())
	fmt.Println("Another value for arithmetic:", reflectVal.Interface().(int))

	// penggunaan fungsi getPropertyInfo()
	s1 := &student{Name: "Atjhoendz", Grade: 100}
	s1.getPropertyInfo()

	// mengambil informasi dari method SetName()
	s2 := &student{Name: "dela", Grade: 90}

	fmt.Println("Nama:", s2.Name)

	reflectValue2 := reflect.ValueOf(s2)
	method := reflectValue2.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("armando"),
	})

	fmt.Println("Nama:", s2.Name)
}
