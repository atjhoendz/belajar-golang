package models

import (
	"fmt"
	"strings"
)

type User struct {
	Name        string
	Email       string
	PhoneNumber *string
	Address     string
}

func (u *User) FormatValue() {
	fmt.Println("----------------[ FormatValue() ]----------------")
	fmt.Println(u)
	u.Name = strings.Title(u.Name)

	u.Email = strings.ToLower(u.Email)

	fmt.Println(u.Address)
	fmt.Println(u.PhoneNumber)
	fmt.Println(*u.PhoneNumber)

	if u.Address != "" {
		fmt.Println(u.Address)
	}

	if u.PhoneNumber != nil {
		*u.PhoneNumber = strings.TrimSpace(*u.PhoneNumber)
	}
	fmt.Println("-------------------[ Batas ] ---------------------")
}

func (u *User) SeeInfo() {
	fmt.Println("----------------[ SeeInfo() ]----------------")
	fmt.Println("u.Name: ", u.Name)
	fmt.Println("u: ", u)
	fmt.Println("*u: ", *u)
	fmt.Println("&u: ", &u)
	fmt.Println("&u.Name: ", &u.Name)
	fmt.Println("u.Name: ", u.Name)
	fmt.Println("-------------------[ Batas ] ---------------------")
}

func (u User) JustSee() {
	fmt.Println(u)
}
