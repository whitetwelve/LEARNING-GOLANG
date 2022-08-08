package main

import (
	"fmt"
)

func main() {
	// ----VARIABEL BOOLEAN---//
	var boolean = true
	fmt.Println(boolean)

	boolean = false
	fmt.Println(boolean)
	// ----VARIABEL INT---//
	var numberic = 100
	fmt.Println(numberic)

	numberic = 2022
	fmt.Println(numberic)
	// // ----VARIABEL STRING---//
	fullname := "Fuad Azkia"
	fmt.Println(fullname)

	fullname = "Ade"
	fmt.Println(fullname)
	// ----DEKLARASIKAN VARIABEL JAMAK----//
	var (
		firstName = "Michael"
		lastName  = "Jackson"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	// // CONST JAMAK //
	const (
		fullName string = "Fuad Azkia"
		age      int    = 30
		address  string = "Bekasi"
	)

	fmt.Println(fullName)
	fmt.Println(age)
	fmt.Println(address)

	//--OPERATOR--//
	var result = 10 + 100
	fmt.Println(result)

	const (
		a int = 200
		b int = 10
		c     = a * b
	)
	fmt.Println(c)
}
