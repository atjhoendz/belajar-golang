package main

import (
	"coba_struct_method/models"
	"fmt"
)

func main() {
	var phoneNumber string = "0918299 99"
	var u = models.User{Name: "moh achun armando", Email: "testEmail@mail.com", PhoneNumber: &phoneNumber}

	fmt.Println(u)

	u.Name = "new Name"

	fmt.Println(u)

	su := &u

	fmt.Println(su)

	su.Name = "another name"

	fmt.Println(*su)

	su.JustSee()

	// u.FormatValue()

	// fmt.Println(u)

	// fmt.Println("---------------")
	// u.SeeInfo()
}
